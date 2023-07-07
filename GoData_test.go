package godata

import (
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"testing"
)

func TestGodata(t *testing.T) {
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

	ds.First()
	for !ds.Eof() {
		t.Log(ds.FieldByName("DESCRICAO").AsString())
		ds.Next()
	}
}

func TestDataSetToStruct(t *testing.T) {

	type Teste struct {
		Descricao *string
	}

	type Process struct {
		QuemAbriu *string
		Lista     []Teste
	}

	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := NewConnection(DialectType(ORACLE), connectStr)
	//db.EnableLog()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	err = ds.
		AddSql("select &total, quem_abriu as QuemAbriu, quem_aprovou as QuemAprovou &from_table").
		AddSql("where cod_empresa between :cod_empresa_ini and :cod_empresa_fim").
		AddSql("and numero_os in (&numero_os)").
		AddSql("and quem_abriu in (&quem_abriu)").
		SetInputParam("cod_empresa_ini", 2).
		SetInputParam("cod_empresa_fim", 35).
		SetMacro("total", "valor_itens_bruto").
		SetMacro("from_table", "from os ").
		SetMacro("numero_os", []int64{2, 100, 23420, 23422, -7}).
		SetMacro("quem_abriu", []string{"LETICIAS", "LEONARDO", "SABRINAP", "MAURILIO"}).
		Open()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ds.FieldByName("QuemAbriu").AsString())

	var dto Process

	err = ds.ToStruct(&dto)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Description: ", *dto.QuemAbriu)
}

func TestDataSetToStructList(t *testing.T) {
	type MedidasDto struct {
		IdMedida      int64   `json:"id_medida"`
		CodMedidaPneu string  `json:"cod_medida_pneu"`
		CodRodaPneu   string  `json:"cod_roda_pneu"`
		Medida        string  `json:"medida"`
		Descricao     string  `json:"descricao"`
		Ativo         string  `json:"ativo"`
		Perimetro     float64 `json:"perimetro"`
	}

	connectStr := "oracle://NBS:NEW@100.0.66.145:1521/NBS"

	db, err := NewConnection(DialectType(ORACLE), connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	err = ds.
		AddSql("SELECT ID_MEDIDA       as idMedida,").
		AddSql("COD_MEDIDA_PNEU as codMedidaPneu,").
		AddSql("COD_RODA_PNEU   as codRodaPneu,").
		AddSql("MEDIDA,").
		AddSql("DESCRICAO,").
		AddSql("PERIMETRO,").
		AddSql("ATIVO").
		AddSql("FROM RECAPAGEM_PNEU_MEDIDA").
		Open()

	if err != nil {
		t.Fatal(err)
	}

	var dto []MedidasDto

	err = ds.ToStruct(&dto)

	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(dto); i++ {
		t.Log(dto[i].IdMedida, dto[i].Descricao)
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

func TestDataSetToSInsertReturn(t *testing.T) {

	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := NewConnectionOracle(connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)

	_, err = ds.
		AddSql("INSERT INTO TESTE (ID_CODIGO_TESTE, DESCRICAO) VALUES (:ID_CODIGO_TESTE, :DESCRICAO)").
		AddSql("RETURNING ID_CODIGO_TESTE, DESCRICAO INTO :OUT_ID_CODIGO_TESTE, :OUT_DESCRICAO").
		SetInputParam("ID_CODIGO_TESTE", 132).
		SetInputParam("DESCRICAO", "INSERT TEST").
		SetOutputParam("OUT_ID_CODIGO_TESTE", int64(0)).
		SetOutputParam("OUT_DESCRICAO", "").
		Exec()

	fmt.Println("ID:", ds.ParamByName("OUT_ID_CODIGO_TESTE").AsInt64())
	fmt.Println("Descrição", ds.ParamByName("OUT_DESCRICAO").AsString())
}

func TestDataSetParseSql(t *testing.T) {

	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := NewConnectionOracle(connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)

	stmt, err := ds.
		AddSql("SELECT").
		AddSql("   A.ID AS ID_ENTIDADE_INTEGRADA").
		AddSql("  ,'N' AS GERAR_ENTIDADE_INTEGRADA").
		AddSql("  ,A.COD_OS_AGENDA").
		AddSql("  ,A.COD_EMPRESA").
		AddSql("  ,A.NUMERO_OS").
		AddSql("  ,CASE").
		AddSql("     WHEN ((A.STATUS NOT IN (1,5) AND B.LIDO = 1 AND B.NUMERO_OS IS NULL)").
		AddSql("             OR (A.STATUS = 0 AND B.LIDO = 0 AND B.NUMERO_OS IS NULL AND B.STATUS_AGENDA = 'C' )) THEN").
		AddSql("       1 --C - Confirmado").
		AddSql("     WHEN (A.STATUS <> 2 AND A.STATUS IN (0,1) AND B.STATUS_AGENDA IN ('A','C') AND (A.DATA_AGENDADA <> B.DATA_AGENDADA)) THEN").
		AddSql("       2 --R - Reagendado").
		AddSql("     WHEN (A.STATUS <> 3 AND C.COD_OS_AGENDA IS NOT NULL) THEN").
		AddSql("       3 --X - Cancelado").
		AddSql("     WHEN ((A.STATUS <> 4) AND B.STATUS_AGENDA IN ('C','F') AND B.NUMERO_OS IS NOT NULL AND B.DATA_ENCERRADA IS NOT NULL) THEN").
		AddSql("       4 --C, F - Serviço Realizado").
		AddSql("     WHEN  (A.STATUS = 0 AND B.STATUS_AGENDA = 'A' AND ((B.DATA_AGENDADA + (1/1440*NVL(D.TOLERANCIA_NO_SHOW,0))) < SYSDATE)) THEN").
		AddSql("       5 -- 'N' - No Show").
		AddSql("     WHEN ((A.STATUS IN (1,2) AND A.NUMERO_OS IS NULL AND B.STATUS_AGENDA = 'C' AND B.NUMERO_OS IS NOT NULL)").
		AddSql("            OR (A.STATUS = 0 AND A.NUMERO_OS IS NULL AND B.STATUS_AGENDA = 'E' AND B.NUMERO_OS IS NOT NULL))	 THEN").
		AddSql("       6 -- C - Serviço Iniciado").
		AddSql("   END AS STATUS").
		AddSql("  ,B.DATA_AGENDADA").
		AddSql("  ,CASE").
		AddSql("     WHEN C.COD_OS_AGENDA IS NULL THEN").
		AddSql("       B.DATA_ULTIMA_ATUALIZACAO").
		AddSql("     ELSE").
		AddSql("       C.DATA_ULTIMA_ATUALIZACAO").
		AddSql("   END DATA_ULTIMA_ATUALIZACAO").
		AddSql("  ,C.DATA_CANCELADA").
		AddSql("FROM FAB_EI_AGD_NISS_SBOK A").
		AddSql("LEFT JOIN OS_AGENDA B ON B.COD_EMPRESA = A.COD_EMPRESA AND B.COD_OS_AGENDA = A.COD_OS_AGENDA").
		AddSql("LEFT JOIN OS_AGENDA_CANC C ON C.COD_EMPRESA = A.COD_EMPRESA AND C.COD_OS_AGENDA = A.COD_OS_AGENDA").
		AddSql("LEFT JOIN PARM_SYS3 D ON D.COD_EMPRESA = A.COD_EMPRESA").
		AddSql("WHERE").
		AddSql("  A.COD_EMPRESA = :COD_EMPRESA").
		AddSql("  AND (").
		AddSql("    --CONFIRMADO").
		AddSql("    (A.STATUS NOT IN (1,5) AND B.LIDO = 1 AND B.NUMERO_OS IS NULL)").
		AddSql("    OR").
		AddSql("    (A.STATUS = 0 AND B.LIDO = 0 AND B.NUMERO_OS IS NULL AND B.STATUS_AGENDA = 'C' )").
		AddSql("    --REAGENDADO").
		AddSql("    OR (A.STATUS <> 2 AND A.STATUS IN (0,1) AND B.STATUS_AGENDA IN ('A','C') AND (A.DATA_AGENDADA <> B.DATA_AGENDADA))").
		AddSql("    --CANCELADO").
		AddSql("    OR (A.STATUS <> 3 AND C.COD_OS_AGENDA IS NOT NULL)").
		AddSql("    --SERVICO REALIZADO").
		AddSql("    OR ((A.STATUS <> 4) AND B.STATUS_AGENDA IN ('C','F') AND B.NUMERO_OS IS NOT NULL AND B.DATA_ENCERRADA IS NOT NULL)").
		AddSql("    --NO SHOW").
		AddSql("    OR (A.STATUS = 0 AND B.STATUS_AGENDA = 'A' AND ((B.DATA_AGENDADA + (1/1440*NVL(D.TOLERANCIA_NO_SHOW,0))) < SYSDATE))").
		AddSql("    --SERVICO INICIADO").
		AddSql("    OR (A.STATUS IN (1,2) AND A.NUMERO_OS IS NULL AND B.STATUS_AGENDA = 'C' AND B.NUMERO_OS IS NOT NULL)").
		AddSql("	--SERVICO INICIADO").
		AddSql("    OR (A.STATUS = 0 AND A.NUMERO_OS IS NULL AND B.STATUS_AGENDA = 'E' AND B.NUMERO_OS IS NOT NULL)").
		AddSql("  )").
		SetInputParam("COD_EMPRESA", 2).
		ParseSql()

	if err != nil {
		t.Fatal(err)
	}

	//st, ok := stmt.(sqlparser.Statement)
	//
	//if ok {
	//	fmt.Println(st)
	//}

	sel, ok := stmt.(*sqlparser.Select)
	if ok {
		for _, expr := range sel.SelectExprs {
			_, ok := expr.(sqlparser.SelectExpr)
			if ok {
				alias, ok := expr.(*sqlparser.AliasedExpr)
				if ok {
					nome, ok := alias.Expr.(*sqlparser.ColName)
					if ok {
						fmt.Println("Column: ", nome.Name)
					}

					if !alias.As.IsEmpty() {
						fmt.Println("Column Alias: ", alias.As)
					}
				}
			}
		}
	}
}
