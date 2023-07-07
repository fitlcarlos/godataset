package godata

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/google/uuid"
	"reflect"
	"strings"
	"unsafe"
)

type DS interface {
	NewDataSet(db *Conn) *DataSet
	Open() error
	Close()
	Exec() (sql.Result, error)
	GetSql() (sql string)
	GetParams() []any
	Scan(list *sql.Rows)
	ParamByName(paramName string) Param
	SetInputParam(paramName string, paramValue any) *DataSet
	SetOutputParam(paramName string, paramType any) *DataSet
	FieldByName(fieldName string) Field
	Locate(key string, value any) bool
	First()
	Next()
	Eof() bool
	IsEmpty() bool
	IsNotEmpty() bool
	Count() int
	AddSql(sql string) *DataSet
	ToStruct(model any) error
}

type DataSet struct {
	Connection      *Conn
	Sql             Strings
	Fields          *Fields
	Rows            []Row
	Params          *Params
	Macros          Macros
	Index           int
	Recno           int
	MasterSouce     *MasterSouce
	IndexFieldNames string
}

func NewDataSet(db *Conn) *DataSet {
	ds := &DataSet{
		Connection:  db,
		Index:       0,
		Recno:       0,
		Fields:      NewFields(),
		Params:      NewParams(),
		Macros:      make(Macros),
		MasterSouce: NewMasterSource(),
	}

	return ds
}

func (ds *DataSet) Open() error {
	ds.Rows = nil
	ds.Index = 0
	ds.Recno = 0

	ds.CreateFields()

	sql := ds.GetSql()

	if ds.Connection.log {
		fmt.Println(sql)
		ds.PrintParam()
	}

	rows, err := ds.Connection.DB.Query(sql, ds.GetParams()...)

	if err != nil {
		errPing := ds.Connection.DB.Ping()
		if errPing != nil {
			errConn := ds.Connection.Open()
			if errConn != nil {
				return err
			}
		} else {
			return fmt.Errorf("could not open dataset %v\n", err)
		}
	}

	defer rows.Close()

	ds.scan(rows)

	ds.First()

	return nil
}

func (ds *DataSet) Close() {
	ds.Sql.Clear()
	ds.Fields = nil
	ds.Rows = nil
	ds.Params = nil
	ds.Macros = nil
	ds.Index = 0
	ds.Recno = 0
	ds.MasterSouce = nil
	ds.IndexFieldNames = ""
}

func (ds *DataSet) Exec() (sql.Result, error) {
	sql := ds.GetSql()

	if ds.Connection.log {
		fmt.Println(sql)
		ds.PrintParam()
	}

	stmt, err := ds.Connection.DB.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.Exec(ds.GetParams()...)
}

func (ds *DataSet) Delete() (int64, error) {
	result, err := ds.Exec()

	if err != nil {
		return 0, err
	}

	rowsAff, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (ds *DataSet) GetSql() (sql string) {

	sql = ds.Sql.Text()

	for key, mrc := range ds.Macros {
		value := reflect.ValueOf(mrc.Value.Value)
		if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
			sql = strings.ReplaceAll(sql, "&"+key, JoinSlice(mrc.Value.AsValue()))
		} else {
			sql = strings.ReplaceAll(sql, "&"+key, mrc.Value.AsString())
		}
	}

	if ds.MasterSouce.DataSource != nil {
		var sqlWhereMasterDetail string

		if len(ds.MasterSouce.MasterFields) != 0 || len(ds.MasterSouce.DetailFields) != 0 {
			for i := 0; i < len(ds.MasterSouce.MasterFields); i++ {

				aliasHash, _ := uuid.NewUUID()
				alias := strings.Replace(aliasHash.String(), "-", "", -1)
				if i == len(ds.MasterSouce.MasterFields)-1 {
					sqlWhereMasterDetail = sqlWhereMasterDetail + ds.MasterSouce.DetailFields[i] + " = :" + alias
				} else {
					sqlWhereMasterDetail = sqlWhereMasterDetail + ds.MasterSouce.DetailFields[i] + " = :" + alias + " and "
				}

				ds.SetInputParam(alias, ds.MasterSouce.DataSource.FieldByName(ds.MasterSouce.MasterFields[i]).AsValue())
			}

			if sqlWhereMasterDetail != "" {
				sql = "select * from (" + sql + ") where " + sqlWhereMasterDetail
			}
		} else {
			fmt.Println("MasterFields or DetailFields field cannot be empty")
		}
	}

	return sql
}

func (ds *DataSet) GetParams() []any {
	var param []any
	for key, prm := range ds.Params.List {
		switch prm.ParamType {
		case IN:
			param = append(param, sql.Named(key, prm.Value.Value))
		case OUT:
			param = append(param, sql.Named(key, sql.Out{Dest: prm.Value.Value, In: false}))
		case INOUT:
			param = append(param, sql.Named(key, sql.Out{Dest: prm.Value.Value, In: true}))
		}
	}
	return param
}

func (ds *DataSet) GetMacros() []any {
	var macro []any
	for key, mrc := range ds.Macros {
		macro = append(macro, sql.Named(key, mrc.Value))
	}
	return macro
}

func (ds *DataSet) scan(list *sql.Rows) {
	fieldTypes, _ := list.ColumnTypes()
	fields, _ := list.Columns()
	for list.Next() {
		columns := make([]any, len(fields))

		for i := range columns {
			columns[i] = &columns[i]
		}

		err := list.Scan(columns...)

		if err != nil {
			print(err)
		}

		row := NewRow()

		for i, value := range columns {
			field := ds.Fields.Add(fields[i])
			field.DataType = fieldTypes[i]
			field.Order = i + 1
			field.Index = i

			row.List[strings.ToUpper(fields[i])] = variant{
				Value: value,
			}
		}

		ds.Rows = append(ds.Rows, row)
	}
}

func (ds *DataSet) ParamByName(paramName string) Param {
	return ds.Params.ParamByName(paramName)
}

func (ds *DataSet) SetInputParam(paramName string, paramValue any) *DataSet {

	ds.Params.List[paramName] = Param{Value: variant{Value: paramValue}, ParamType: IN}

	return ds
}

func (ds *DataSet) SetOutputParam(paramName string, paramType any) *DataSet {
	switch paramType.(type) {
	case int, int8, int16, int32, int64:
		dataType := int64(0)
		ds.Params.List[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case float32:
		dataType := float32(0)
		ds.Params.List[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case float64:
		dataType := float64(0)
		ds.Params.List[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case string:
		dataType := generateString()
		ds.Params.List[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	default:
		dataType := float64(0)
		ds.Params.List[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	}
	return ds
}

func (ds *DataSet) SetMacro(macroName string, macroValue any) *DataSet {

	ds.Macros[macroName] = Macro{Value: variant{Value: macroValue}}

	return ds
}

func (ds *DataSet) CreateFields() error {

	stmt, err := sqlparser.Parse(ds.GetSql())

	if err != nil {
		return err
	}

	ds.Fields = NewFields()
	ds.Fields.Owner = ds

	sel, ok := stmt.(*sqlparser.Select)
	if ok {
		for _, expr := range sel.SelectExprs {
			_, ok := expr.(sqlparser.SelectExpr)
			if ok {
				alias, ok := expr.(*sqlparser.AliasedExpr)
				if ok {
					if !alias.As.IsEmpty() {
						_ = ds.Fields.Add(alias.As.String())
					} else {
						column, ok := alias.Expr.(*sqlparser.ColName)
						if ok {
							_ = ds.Fields.Add(column.Name.String())
						}
					}
				}
			}
		}
	}

	return nil
}

func (ds *DataSet) FieldByName(fieldName string) *Field {
	return ds.Fields.FieldByName(fieldName)
}

func (ds *DataSet) Locate(key string, value any) bool {

	ds.First()
	for ds.Eof() == false {
		switch value.(type) {
		case string:
			if ds.FieldByName(key).AsValue() == value {
				return true
			}
		default:
			if ds.FieldByName(key).AsValue() == value {
				return true
			}
		}

		ds.Next()
	}
	return false
}

func (ds *DataSet) First() {
	ds.Index = 0
	ds.Recno = 0
	if ds.Count() > 0 {
		ds.Recno = 1
	}
}

func (ds *DataSet) Next() {
	if !ds.Eof() {
		ds.Index++
		ds.Recno++
	}
}

func (ds *DataSet) Eof() bool {
	return ds.Count() == 0 || ds.Recno > ds.Count()
}

func (ds *DataSet) IsEmpty() bool {
	return ds.Count() == 0
}

func (ds *DataSet) IsNotEmpty() bool {
	return ds.Count() > 0
}

func (ds *DataSet) Count() int {
	return len(ds.Rows)
}

func (ds *DataSet) AddSql(sql string) *DataSet {
	ds.Sql.Add(sql)

	return ds
}

func (ds *DataSet) AddMasterSource(dataSet *DataSet) *DataSet {
	ds.MasterSouce.AddMasterSource(dataSet)
	return ds
}

func (ds *DataSet) AddDetailFields(fields ...string) *DataSet {
	ds.MasterSouce.AddDetailFields(fields...)
	return ds
}

func (ds *DataSet) AddMasterFields(fields ...string) *DataSet {
	ds.MasterSouce.AddMasterFields(fields...)
	return ds
}

func (ds *DataSet) ToStruct(model any) error {
	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)

	if modelType.Kind() == reflect.Pointer {
		modelType = modelValue.Type().Elem()
		modelValue = modelValue.Elem()
	} else {
		return fmt.Errorf("the variable " + modelType.Name() + " is not a pointer.")
	}

	switch modelType.Kind() {
	case reflect.Struct:
		return ds.toStructUniqResult(modelValue)
	case reflect.Slice, reflect.Array:
		return ds.toStructList(modelValue)
	default:
		return errors.New("The interface is not a slice, array or struct")
	}
}

func (ds *DataSet) toStructUniqResult(modelValue reflect.Value) error {
	for i := 0; i < modelValue.NumField(); i++ {
		if modelValue.Field(i).Kind() != reflect.Slice && modelValue.Field(i).Kind() != reflect.Array {
			field := modelValue.Type().Field(i)
			fieldName := field.Tag.Get("column")

			if fieldName == "" {
				fieldName = field.Name
			}

			if field.Anonymous {
				continue
			}

			var fieldValue any

			if field.Type.Kind() == reflect.Pointer {
				if ds.FieldByName(fieldName).IsNotNull() {
					fieldValue = ds.GetValue(ds.FieldByName(fieldName), modelValue.Field(i).Interface())

					if fieldValue != nil {
						rf := modelValue.Field(i)
						rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
						rf.Set(reflect.ValueOf(fieldValue))
					}
				}
			} else {
				fieldValue = ds.GetValue(ds.FieldByName(fieldName), modelValue.Field(i).Interface())

				if fieldValue != nil {
					rf := modelValue.Field(i)
					rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
					rf.Set(reflect.ValueOf(fieldValue))
				}
			}
		}
	}

	return nil
}

func (ds *DataSet) GetValue(field *Field, fieldType any) any {
	var fieldValue any
	valueType := reflect.TypeOf(fieldType)
	switch fieldType.(type) {
	case int:
		val := field.AsInt()
		fieldValue = reflect.ValueOf(val).Convert(valueType).Interface()
	case int8:
		val := field.AsInt8()
		fieldValue = reflect.ValueOf(val).Convert(valueType).Interface()
	case int16:
		val := field.AsInt16()
		fieldValue = reflect.ValueOf(val).Convert(valueType).Interface()
	case int32:
		val := field.AsInt32()
		fieldValue = reflect.ValueOf(val).Convert(valueType).Interface()
	case int64:
		val := field.AsInt64()
		fieldValue = reflect.ValueOf(val).Convert(valueType).Interface()
	case float32, float64:
		val := field.AsFloat64()
		fieldValue = reflect.ValueOf(val).Convert(valueType).Interface()
	case string:
		val := field.AsString()
		fieldValue = reflect.ValueOf(val).Convert(valueType).Interface()
	case *int:
		val := field.AsInt()
		fieldValue = reflect.ValueOf(&val).Convert(valueType).Interface()
	case *int8:
		val := field.AsInt8()
		fieldValue = reflect.ValueOf(&val).Convert(valueType).Interface()
	case *int16:
		val := field.AsInt16()
		fieldValue = reflect.ValueOf(&val).Convert(valueType).Interface()
	case *int32:
		val := field.AsInt32()
		fieldValue = reflect.ValueOf(&val).Convert(valueType).Interface()
	case *int64:
		val := field.AsInt64()
		fieldValue = reflect.ValueOf(&val).Convert(valueType).Interface()
	case *float32, *float64:
		val := field.AsFloat64()
		fieldValue = reflect.ValueOf(&val).Convert(valueType).Interface()
	case *string:
		val := field.AsString()
		fieldValue = reflect.ValueOf(&val).Convert(valueType).Interface()
	default:
		fieldValue = field.AsValue()
	}
	return fieldValue
}

func (ds *DataSet) toStructList(modelValue reflect.Value) error {
	var modelType reflect.Type

	if modelValue.Type().Kind() == reflect.Pointer {
		modelType = modelValue.Type().Elem().Elem()
		modelValue = modelValue.Elem()
	} else {
		modelType = modelValue.Type().Elem()
	}

	for !ds.Eof() {
		var newModel reflect.Value

		newModel = reflect.New(modelType)

		err := ds.toStructUniqResult(reflect.ValueOf(newModel.Interface()).Elem())

		if err != nil {
			return err
		}

		modelValue.Set(reflect.Append(modelValue, newModel.Elem()))

		ds.Next()
	}

	return nil
}

func JoinSlice(list any) string {
	valueOf := reflect.ValueOf(list)
	value := make([]string, valueOf.Len())
	for i := 0; i < valueOf.Len(); i++ {
		if valueOf.Index(i).Type().Kind() == reflect.String {
			value[i] = "'" + fmt.Sprintf("%v", valueOf.Index(i).Interface()) + "'"
		} else {
			value[i] = fmt.Sprintf("%v", valueOf.Index(i).Interface())
		}
	}
	return strings.Join(value, ", ")
}

func (ds *DataSet) PrintParam() {
	for key, value := range ds.Params.List {
		fmt.Println("Colum:", key, "Value:", value.AsValue(), "Type:", reflect.TypeOf(value.AsValue()))
	}
}

func generateString() string {
	var result string
	for i := 0; i < 400; i++ {
		result += "abcdefghij"
	}
	return result
}

func (ds *DataSet) ParseSql() (sqlparser.Statement, error) {
	return sqlparser.Parse(ds.GetSql())
}
