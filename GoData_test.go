package godata

import "testing"

func TestGodata(t *testing.T) {
	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := NewConnection(DialectType(ORACLE), connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	ds.
		AddSql("SELECT DESCRICAO FROM FAB_PROCESSO").
	    AddSql("WHERE ID BETWEEN :idini AND :idfim").
		SetParam("idini", 20).
		SetParam("idfim", 100).
		Open()

	ds.First()
	for !ds.Eof() {
		t.Log(ds.FieldByName("DESCRICAO").AsString())
		ds.Next()
	}
}

func TestDataSetToStruct(t *testing.T) {
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
	ds.
		AddSql("SELECT DESCRICAO FROM FAB_PROCESSO").
		AddSql("WHERE ID BETWEEN :idini AND :idfim").
		SetParam("idini", 20).
		SetParam("idfim", 100).
		Open()

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
	ds.
		AddSql("SELECT DESCRICAO FROM FAB_PROCESSO").
		AddSql("WHERE ID BETWEEN :idini AND :idfim").
		SetParam("idini", 20).
		SetParam("idfim", 100).
		Open()

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
	_,_ = ds.
			AddSql("INSERT INTO TESTE (ID_CODIGO_TESTE) VALUES (:ID_CODIGO_TESTE)").
			SetParam("ID_CODIGO_TESTE", 100).
			Exec()
}