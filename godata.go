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
	Param            Params
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
		Param:      make(Params),
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
	ds.Param   = nil
	ds.Index = 0
	ds.Recno = 0
	ds.DetailFields = ""
	ds.MasterSouce = nil
	ds.MasterFields = ""
	ds.MasterDetailList = nil
	ds.IndexFieldNames = ""
}

func (ds *DataSet) Exec() error {

	result, err := ds.Connection.DB.Exec(ds.GetSql(), ds.GetParams()...)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("error to execute query: %w", err)
	}

	return nil
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

			ds.ParamByName(alias, ds.MasterSouce.FieldByName(mf[i]).Value)
		}

		if sqlWhereMasterDetail != "" {
			sql = "select * from (" + sql + ") where " + sqlWhereMasterDetail
		}
	}

	return sql
}

func (ds *DataSet) GetParams() []any {
	var param []any
	for _, prm := range ds.Param {
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

func (ds *DataSet) ParamByName(paramName string, paramValue any) *DataSet {

	ds.Param[paramName] = Parameter{Value: paramValue}

	return ds
}

func (ds *DataSet) FieldByName(fieldName string) Field {
	fieldName = strings.ToUpper(fieldName)

	return ds.Rows[ds.Index][fieldName]
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
		ds.ToStructUniqResult(model)
	case reflect.Slice, reflect.Array:
		ds.ToStructList(model)
	default:
		return errors.New("The interface is not a slice, array or struct")
	}
	return nil
}

func (ds *DataSet) ToStructUniqResult(model any) error {
	modelType := reflect.TypeOf(model)
	modelValue:= reflect.ValueOf(model)

	if modelType.Kind() == reflect.Pointer {
		modelType  = reflect.TypeOf(model).Elem()
		modelValue = reflect.ValueOf(model).Elem()
	}

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		if field.Anonymous {
			continue
		}

		fieldValue := reflect.New(field.Type).Elem().Interface()
		fieldType := reflect.New(field.Type).Elem().Type()

		switch fieldValue.(type) {
		case float32, float64:
			val := ds.FieldByName(field.Name).AsFloat64()
			fieldValue = reflect.ValueOf(val).Convert(fieldType).Interface()
		default:
			fieldValue = ds.FieldByName(field.Name).AsValue()
		}

		rf := modelValue.Field(i)
		rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
		rf.Set(reflect.ValueOf(fieldValue))
	}

	return nil
}

func (ds *DataSet) ToStructList(model any) error {
	modelType := reflect.TypeOf(model).Elem().Elem()
	modelValue := reflect.ValueOf(model).Elem()

	for !ds.Eof() {
		var newModel reflect.Value

		if modelType.Kind() == reflect.Ptr {
			newModel = reflect.New(modelType.Elem()).Elem()
		}else{
			newModel = reflect.New(modelType)
		}

		ds.ToStructUniqResult(newModel.Interface())

		modelValue.Set(reflect.Append(modelValue, newModel.Elem()))

		ds.Next()
	}

	return nil
}

