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
		AddSql("SELECT DESCRICAO FROM FAB_PROCESSO WHERE ID BETWEEN :idini AND :idfim").
		ParamByName("idini", 20).
		ParamByName("idfim", 100).
		Open()

	ds.First()
	for !ds.Eof() {
		t.Log(ds.FieldByName("DESCRICAO").AsString())
		ds.Next()
	}
}
