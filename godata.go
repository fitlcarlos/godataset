package godata

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
	"unsafe"
)

type DataSet struct {
	Connection       *Conn
	Sql              Strings
	Columns          []string
	Rows             Fields
	Params           Params
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
		Params:     make(Params),
	}

	return ds
}

func (ds *DataSet) Open() error {
	ds.Rows = nil
	ds.Index = 0
	ds.Recno = 0

	rows, err := ds.Connection.DB.Query(ds.GetSql(), ds.GetParams()...)

	if err != nil {
		return fmt.Errorf("could not open dataset %v\n",err)
	}

	col, _ := rows.Columns()
	ds.Columns = col

	defer rows.Close()

	ds.Scan(rows)

	ds.First()

	return nil
}

func (ds *DataSet) Close() {
	ds.Columns = nil
	ds.Rows    = nil
	ds.Params  = nil
	ds.Index = 0
	ds.Recno = 0
	ds.DetailFields = ""
	ds.MasterSouce = nil
	ds.MasterFields = ""
	ds.MasterDetailList = nil
	ds.IndexFieldNames = ""
}

func (ds *DataSet) Exec() (sql.Result, error) {
	return ds.Connection.DB.Exec(ds.GetSql(), ds.GetParams()...)
}

func (ds *DataSet) GetSql() (sql string) {

	sql = ds.Sql.Text()

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

			ds.SetParam(alias, ds.MasterSouce.FieldByName(mf[i]).Value)
		}

		if sqlWhereMasterDetail != "" {
			sql = "select * from (" + sql + ") where " + sqlWhereMasterDetail
		}
	}

	return sql
}

func (ds *DataSet) GetParams() []any {
	var param []any
	for _, prm := range ds.Params {
		param = append(param, prm.Value)
	}
	return param
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
			row[fields[i]] = Field{
				Name:       fields[i],
				Caption:    fields[i],
				DataType:   columntypes[i],
				Value:      variant{Value: value},
				DataMask:   "",
				ValueTrue:  "",
				ValueFalse: "",
				Visible:    true,
				Order:      i + 1,
				Index:      i,
			}
		}

		ds.Rows = append(ds.Rows, row)
	}
}

func (ds *DataSet) ParamByName(paramName string) Param {
	return ds.Params[paramName]
}

func (ds *DataSet) SetParam(paramName string, paramValue any) *DataSet {

	ds.Params[paramName] = Param{Value: paramValue}

	return ds
}

func (ds *DataSet) FieldByName(fieldName string) Field {
	fieldName = strings.ToUpper(fieldName)

	if len(ds.Rows) > 0{
		return ds.Rows[ds.Index][fieldName]
	}else{
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
	if !ds.Eof(){
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

	switch reflect.TypeOf(model).Elem().Kind() {
	case reflect.Struct:
		ds.toStructUniqResult(model)
	case reflect.Slice, reflect.Array:
		ds.toStructList(model)
	default:
		return errors.New("The interface is not a slice, array or struct")
	}
	return nil
}

func (ds *DataSet) toStructUniqResult(model any) error {
	modelType := reflect.TypeOf(model)
	modelValue:= reflect.ValueOf(model)

	if modelType.Kind() == reflect.Pointer {
		modelType  = reflect.TypeOf(model).Elem()
		modelValue = reflect.ValueOf(model).Elem()
	}

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
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

		rf := modelValue.Field(i)
		rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
		rf.Set(reflect.ValueOf(fieldValue))
	}

	return nil
}

func (ds *DataSet) toStructList(model any) error {
	modelType := reflect.TypeOf(model).Elem().Elem()
	modelValue := reflect.ValueOf(model).Elem()

	for !ds.Eof() {
		var newModel reflect.Value

		if modelType.Kind() == reflect.Ptr {
			newModel = reflect.New(modelType.Elem()).Elem()
		}else{
			newModel = reflect.New(modelType)
		}

		ds.toStructUniqResult(newModel.Interface())

		modelValue.Set(reflect.Append(modelValue, newModel.Elem()))

		ds.Next()
	}

	return nil
}

