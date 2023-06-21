package godata

import (
	"fmt"
	"testing"
)

//func TestGodata(t *testing.T) {
//	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
//
//	db, err := NewConnection(DialectType(ORACLE), connectStr)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	defer db.Close()
//
//	ds := NewDataSet(db)
//	err = ds.
//		AddSql("SELECT DESCRICAO FROM FAB_PROCESSO").
//		AddSql("WHERE ID BETWEEN :idini AND :idfim").
//		SetInputParam("idini", 20).
//		SetInputParam("idfim", 100).
//		Open()
//
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	ds.First()
//	for !ds.Eof() {
//		t.Log(ds.FieldByName("DESCRICAO").AsString())
//		ds.Next()
//	}
//}

func TestDataSetToStruct(t *testing.T) {

	type Teste struct {
		Descricao string
	}

	type Process struct {
		Descricao string
		Lista     []Teste
	}

	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := NewConnection(DialectType(ORACLE), connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	err = ds.
		AddSql("select &total &from_table").
		AddSql("where cod_empresa between :cod_empresa_ini and :cod_empresa_fim").
		AddSql("and numero_os in (&numero_os)").
		SetInputParam("cod_empresa_ini", 2).
		SetInputParam("cod_empresa_fim", 33).
		SetMacro("total", "valor_itens_bruto").
		SetMacro("from_table", "from os ").
		SetMacro("numero_os", []int64{2, 100, 23420, 23422}).
		Open()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Sprintln(ds.FieldByName("valor_itens_bruto").AsString())

	var dto Process

	err = ds.ToStruct(&dto)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(dto.Descricao)
}

func TestDataSetToStructList(t *testing.T) {
	type Process struct {
		Descricao string
	}

	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := NewConnection(DialectType(ORACLE), connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	err = ds.
		AddSql("SELECT DESCRICAO FROM FAB_PROCESSO").
		AddSql("WHERE ID BETWEEN :idini AND :idfim").
		SetInputParam("idini", 20).
		SetInputParam("idfim", 100).
		Open()

	if err != nil {
		t.Fatal(err)
	}

	var dto []Process

	err = ds.ToStruct(&dto)

	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(dto); i++ {
		t.Log(dto[i].Descricao)
	}
}

func TestDataSetToSInsert(t *testing.T) {

	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := NewConnection(DialectType(ORACLE), connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	_, _ = ds.
		AddSql("INSERT INTO TESTE (ID_CODIGO_TESTE) VALUES (:ID_CODIGO_TESTE)").
		SetInputParam("ID_CODIGO_TESTE", 100).
		Exec()
}
