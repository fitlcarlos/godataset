# godata

GoData is a library that facilitates interaction with a database and manipulation of data through a dataset.

A dataset consists of a series of records, each containing any number of fields and a pointer to the current record. The dataset may have a direct, one-to-one correspondence with a physical table, or, as a result of a query, it may be a subset of a table or a join of several tables.

Example of use:

  connectStr := "oracle://erp:pass123@DESKTOP-DEV:1521/xe"

	db, err := NewConnection(DialectType(ORACLE), connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	err = ds.
		AddSql("SELECT ID, NAME FROM PEOPLE").
		AddSql("WHERE ID BETWEEN :idStart AND :idEnd").
		SetInputParam("idStart", 20).
		SetInputParam("idEnd", 100).
		Open()

	if err != nil {
		t.Fatal(err)
	}

  fmt.Println(ds.Count())

	ds.First()
	for !ds.Eof() {
		t.Log(ds.FieldByName("NAME").AsString())
		ds.Next()
	}

All DataSet methods 

func NewDataSet(db *Conn) *DataSet
func NewDataSetTx(tx *Transaction) *DataSet
func (ds *DataSet) AddContext(ctx context.Context) *DataSet
func (ds *DataSet) Open() error
func (ds *DataSet) OpenContext(context context.Context) error
func (ds *DataSet) Close()
func (ds *DataSet) Exec() (sql.Result, error)
func (ds *DataSet) ExecContext(context context.Context) (sql.Result, error)
func (ds *DataSet) Delete() (int64, error)
func (ds *DataSet) DeleteContext(context context.Context) (int64, error)
func (ds *DataSet) GetSql() (sql string)
func (ds *DataSet) GetSqlMasterDetail() (vsql string)
func (ds *DataSet) GetParams() []any
func (ds *DataSet) GetMacros() []any
func (ds *DataSet) ParamByName(paramName string) *Param
func (ds *DataSet) MacroByName(macroName string) *Macro
func (ds *DataSet) SetInputParam(paramName string, paramValue any) *DataSet
func (ds *DataSet) SetInputParamClob(paramName string, paramValue string) *DataSet
func (ds *DataSet) SetInputParamBlob(paramName string, paramValue []byte) *DataSet
func (ds *DataSet) SetOutputParam(paramName string, paramValue any) *DataSet
func (ds *DataSet) SetOutputParam(paramName string, paramValue any) *DataSet
func (ds *DataSet) SetMacro(macroName string, macroValue any) *DataSet
func (ds *DataSet) CreateFields() error
func (ds *DataSet) Prepare()
func (ds *DataSet) FieldByName(fieldName string) *Field
func (ds *DataSet) Locate(key string, value any) bool
func (ds *DataSet) First()
func (ds *DataSet) Next()
func (ds *DataSet) Previous()
func (ds *DataSet) Last()
func (ds *DataSet) Bof() bool
func (ds *DataSet) Eof() bool
func (ds *DataSet) IsEmpty() bool
func (ds *DataSet) IsNotEmpty() bool
func (ds *DataSet) Count() int
func (ds *DataSet) AddSql(sql string) *DataSet
func (ds *DataSet) AddMasterSource(dataSet *DataSet) *DataSet 
func (ds *DataSet) AddDetailFields(fields ...string) *DataSet
func (ds *DataSet) AddMasterFields(fields ...string) *DataSet
func (ds *DataSet) ClearDetailFields() *DataSet 
func (ds *DataSet) ClearMasterFields() *DataSet
func (ds *DataSet) ToStruct(model any) error 
func (ds *DataSet) toStructUniqResult(modelValue reflect.Value) error
func (ds *DataSet) GetValue(field *Field, fieldType any) any
func (ds *DataSet) PrintParam()
func (ds *DataSet) ParseSql() (sqlparser.Statement, error)
func (ds *DataSet) SqlParam() string


