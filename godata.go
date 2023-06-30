package godata

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
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
	Connection       *Conn
	Sql              Strings
	Columns          []string
	Rows             Fields
	Params           *Params
	Macros           Macros
	Index            int
	Recno            int
	DetailFields     string
	MasterSouce      *DataSet
	MasterFields     string
	MasterDetailList map[string]MasterDetails
	IndexFieldNames  string
}

func NewDataSet(db *Conn) *DataSet {
	ds := &DataSet{
		Connection: db,
		Index:      0,
		Recno:      0,
		Params:     NewParams(),
		Macros:     make(Macros),
	}

	return ds
}

func (ds *DataSet) Open() error {
	ds.Rows = nil
	ds.Index = 0
	ds.Recno = 0

	sql := ds.GetSql()

	if ds.Connection.log {
		log.Println(sql)
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

	col, _ := rows.Columns()
	ds.Columns = col

	defer rows.Close()

	ds.Scan(rows)

	ds.First()

	return nil
}

func (ds *DataSet) Close() {
	ds.Sql.Clear()
	ds.Columns = nil
	ds.Rows = nil
	ds.Params = nil
	ds.Macros = nil
	ds.Index = 0
	ds.Recno = 0
	ds.DetailFields = ""
	ds.MasterSouce = nil
	ds.MasterFields = ""
	ds.MasterDetailList = nil
	ds.IndexFieldNames = ""
}

func (ds *DataSet) Exec() (sql.Result, error) {
	sql := ds.GetSql()

	if ds.Connection.log {
		log.Println(sql)
		ds.PrintParam()
	}

	stmt, err := ds.Connection.DB.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.Exec(ds.GetParams()...)
}

//func (ds *DataSet) InsertReturning(fieldName ...output) (*Params, error) {
//
//	vsql := ds.GetSql()
//
//	if ds.Connection.log {
//		log.Println(vsql)
//		ds.PrintParam()
//	}
//
//	//joinFieldName := strings.Join(fieldName, ",")
//
//	var params []any
//	for key, param := range ds.Params.value {
//		switch param.ParamType {
//		case IN:
//			params = append(params, sql.Named(key, param.Value.Value))
//		}
//	}
//
//	var joinFieldName string
//	var paramStr string
//	var vparam *Params = NewParams()
//
//	count := 0
//	//for i := 0; i < len(fieldName); i++
//	for key, value := range fieldName {
//		if count < len(fieldName)-1 {
//			joinFieldName = joinFieldName + key + ","
//			paramStr = paramStr + ":OUT_" + strings.ToUpper(key) + ","
//		} else if count == len(fieldName)-1 {
//			joinFieldName = joinFieldName + key
//			paramStr = paramStr + ":OUT_" + strings.ToUpper(key)
//		}
//
//		param := NewParam(OUT)
//		vparam.value[key] = param
//		out := &param.Value.Value
//
//		//params = append(params, sql.Named("OUT_"+strings.ToUpper(fieldName[i]), sql.Out{Dest: out}))
//		params = append(params, sql.Out{Dest: out})
//
//		count++
//	}
//
//	var returningStr string
//	switch ds.Connection.Dialect {
//	case ORACLE:
//		returningStr = returningStr + " RETURNING " + joinFieldName + " INTO "
//	case POSTGRESQL:
//		returningStr = returningStr + strings.ToUpper(" RETURNING "+joinFieldName)
//	}
//
//	stmt, err := ds.Connection.DB.Prepare(vsql + returningStr + paramStr)
//
//	if err != nil {
//		return nil, err
//	}
//
//	defer stmt.Close()
//
//	_, err = stmt.Exec(params)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return vparam, nil
//}

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

	if ds.MasterSouce != nil {
		var sqlWhereMasterDetail string
		mf := strings.Split(ds.MasterFields, ";")
		df := strings.Split(ds.DetailFields, ";")

		for i := 0; i < len(mf); i++ {
			aliasHash, _ := uuid.NewUUID()
			alias := strings.Replace(aliasHash.String(), "-", "", -1)
			if i == len(mf)-1 {
				sqlWhereMasterDetail = sqlWhereMasterDetail + df[i] + " = :" + alias
			} else {
				sqlWhereMasterDetail = sqlWhereMasterDetail + df[i] + " = :" + alias + " and "
			}

			ds.SetInputParam(alias, ds.MasterSouce.FieldByName(mf[i]).Value)
		}

		if sqlWhereMasterDetail != "" {
			sql = "select * from (" + sql + ") where " + sqlWhereMasterDetail
		}
	}

	return sql
}

func (ds *DataSet) GetParams() []any {
	var param []any
	for key, prm := range ds.Params.value {
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

func (ds *DataSet) Scan(list *sql.Rows) {
	columntypes, _ := list.ColumnTypes()
	fields, _ := list.Columns()
	for list.Next() {
		columns := make([]interface{}, len(fields))

		for i := range columns {
			columns[i] = &columns[i]
		}

		err := list.Scan(columns...)

		if err != nil {
			print(err)
		}

		row := make(map[string]Field)

		for i, value := range columns {
			field := NewField(fields[i], columntypes[i])
			field.Value.Value = value
			field.Order = i + 1
			field.Index = i

			row[fields[i]] = field
		}

		ds.Rows = append(ds.Rows, row)
	}
}

func (ds *DataSet) ParamByName(paramName string) Param {
	return ds.Params.ParamByName(paramName)
}

func (ds *DataSet) SetInputParam(paramName string, paramValue any) *DataSet {

	ds.Params.value[paramName] = Param{Value: variant{Value: paramValue}, ParamType: IN}

	return ds
}

func (ds *DataSet) SetOutputParam(paramName string, paramType any) *DataSet {
	switch paramType.(type) {
	case int, int8, int16, int32, int64:
		dataType := int64(0)
		ds.Params.value[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case float32:
		dataType := float32(0)
		ds.Params.value[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case float64:
		dataType := float64(0)
		ds.Params.value[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case string:
		dataType := generateString()
		ds.Params.value[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	default:
		dataType := float64(0)
		ds.Params.value[paramName] = Param{Value: variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	}
	return ds
}

func (ds *DataSet) SetMacro(macroName string, macroValue any) *DataSet {

	ds.Macros[macroName] = Macro{Value: variant{Value: macroValue}}

	return ds
}

func (ds *DataSet) FieldByName(fieldName string) Field {
	fieldName = strings.ToUpper(fieldName)

	if len(ds.Rows) > 0 {
		return ds.Rows[ds.Index][fieldName]
	} else {
		return Field{}
	}
}

func (ds *DataSet) Locate(key string, value any) bool {

	ds.First()
	for ds.Eof() == false {
		switch value.(type) {
		case string:
			if ds.FieldByName(key).Value == value {
				return true
			}
		default:
			if ds.FieldByName(key).Value == value {
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

			fieldValue := reflect.New(field.Type).Elem().Interface()
			fieldType := reflect.New(field.Type).Elem().Type()

			switch fieldValue.(type) {
			case int:
				val := ds.FieldByName(fieldName).AsInt()
				fieldValue = reflect.ValueOf(val).Convert(fieldType).Interface()
			case int8:
				val := ds.FieldByName(fieldName).AsInt8()
				fieldValue = reflect.ValueOf(val).Convert(fieldType).Interface()
			case int16:
				val := ds.FieldByName(fieldName).AsInt16()
				fieldValue = reflect.ValueOf(val).Convert(fieldType).Interface()
			case int32:
				val := ds.FieldByName(fieldName).AsInt32()
				fieldValue = reflect.ValueOf(val).Convert(fieldType).Interface()
			case int64:
				val := ds.FieldByName(fieldName).AsInt64()
				fieldValue = reflect.ValueOf(val).Convert(fieldType).Interface()
			case float32, float64:
				val := ds.FieldByName(fieldName).AsFloat64()
				fieldValue = reflect.ValueOf(val).Convert(fieldType).Interface()
			default:
				fieldValue = ds.FieldByName(fieldName).AsValue()
			}

			if fieldValue != nil {
				rf := modelValue.Field(i)
				rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
				rf.Set(reflect.ValueOf(fieldValue))
			}
		}
	}

	return nil
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
	for key, value := range ds.Params.value {
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
