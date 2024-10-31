package godata

import (
	"fmt"
	"testing"
)

func TestGodata(t *testing.T) {
	connectStr := "oracle://erp:100651xpto@DESKTOP-AU8VNS3:1521/xe"

	db, err := NewConnection(DialectType(ORACLE), connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	err = ds.
		AddSql("SELECT id, nome FROM pessoa").
		AddSql("WHERE ID BETWEEN :idini AND :idfim").
		SetInputParam("idini", 20).
		SetInputParam("idfim", 100).
		Open()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ds.Count())

	ds.First()
	for !ds.Eof() {
		t.Log(ds.FieldByName("nome").AsString())
		ds.Next()
	}

	ds.Close()

	//t.Log(ds.FieldByName("UCID").AsString())
}

func TestDataSetToStruct(t *testing.T) {

	t.Log("Sucesso.")

	//type Teste struct {
	//	Descricao *string
	//}
	//
	//type Process struct {
	//	QuemAbriu *string
	//	Lista     []Teste
	//}
	//
	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnection(DialectType(ORACLE), connectStr)
	//db.PoolSize = 20
	////db.EnableLog()
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//ds := NewDataSet(db)
	//err = ds.
	//	AddSql("select &total, quem_abriu as QuemAbriu, quem_aprovou as QuemAprovou &from_table").
	//	AddSql("where cod_empresa between :cod_empresa_ini and :cod_empresa_fim").
	//	AddSql("and numero_os in (&numero_os)").
	//	AddSql("and quem_abriu in (&quem_abriu)").
	//	SetInputParam("cod_empresa_ini", 2).
	//	SetInputParam("cod_empresa_fim", 35).
	//	SetMacro("total", "valor_itens_bruto").
	//	SetMacro("from_table", "from os ").
	//	SetMacro("numero_os", []int64{2, 100, 23420, 23422, -7}).
	//	SetMacro("quem_abriu", []string{"LETICIAS", "LEONARDO", "SABRINAP", "MAURILIO"}).
	//	Open()
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//fmt.Println(ds.FieldByName("QuemAbriu").AsString())
	//fmt.Println(ds.ParamByName("cod_empresa_ini").AsString())
	//fmt.Println(ds.MacroByName("total").AsString())
	//
	//var dto Process
	//
	//err = ds.ToStruct(&dto)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//t.Log("Description: ", *dto.QuemAbriu)
}

func TestDataSetToStructList(t *testing.T) {

	t.Log("Sucesso.")

	//type MedidasDto struct {
	//	IdMedida      int64   `json:"id_medida"`
	//	CodMedidaPneu string  `json:"cod_medida_pneu"`
	//	CodRodaPneu   string  `json:"cod_roda_pneu"`
	//	Medida        string  `json:"medida"`
	//	Descricao     string  `json:"descricao"`
	//	Ativo         string  `json:"ativo"`
	//	Perimetro     float64 `json:"perimetro"`
	//}
	//
	//connectStr := "oracle://NBS:NEW@100.0.66.145:1521/NBS"
	//
	//db, err := NewConnection(DialectType(ORACLE), connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//ds := NewDataSet(db)
	//err = ds.
	//	AddSql("SELECT ID_MEDIDA       as idMedida,").
	//	AddSql("COD_MEDIDA_PNEU as codMedidaPneu,").
	//	AddSql("COD_RODA_PNEU   as codRodaPneu,").
	//	AddSql("MEDIDA,").
	//	AddSql("DESCRICAO,").
	//	AddSql("PERIMETRO,").
	//	AddSql("ATIVO").
	//	AddSql("FROM RECAPAGEM_PNEU_MEDIDA").
	//	Open()
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//var dto []MedidasDto
	//
	//err = ds.ToStruct(&dto)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//for i := 0; i < len(dto); i++ {
	//	t.Log(dto[i].IdMedida, dto[i].Descricao)
	//}
}

func TestDataSetToSInsert(t *testing.T) {

	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnection(DialectType(ORACLE), connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//ds := NewDataSet(db)
	//
	//_, err = ds.
	//	AddSql("INSERT INTO TESTE (ID_CODIGO_TESTE, DESCRICAO)").
	//	AddSql("VALUES (:ID_CODIGO_TESTE, :DESCRICAO)").
	//	AddSql("RETURNING DESCRICAO into :out_desc").
	//	SetInputParam("ID_CODIGO_TESTE", 182).
	//	SetInputParam("DESCRICAO", "Testesssssssssssss").
	//	SetOutputParam("out_desc", string("")).
	//	Exec()
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//fmt.Println(ds.ParamByName("desc").AsString())
	//fmt.Println(ds.FieldByName("desc").AsString())
	//fmt.Println(ds.MacroByName("desc").AsString())
}

func TestDataSetToSInsertReturn(t *testing.T) {

	t.Log("Sucesso.")

	// connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	// db, err := NewConnectionOracle(connectStr)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// defer db.Close()

	// ds := NewDataSet(db)

	// _, err = ds.
	// 	AddSql("INSERT INTO TESTE (ID_CODIGO_TESTE, DESCRICAO) VALUES (161, 'INSERT TEST')").
	// 	AddSql("RETURNING ID_CODIGO_TESTE INTO :ID").
	// 	//SetInputParam("ABC", 1).
	// 	//SetInputParam("DESCRICAO", "INSERT TEST").
	// 	SetOutputParam("ID", int64(0)).
	// 	//SetOutputParam("OUT_DESCRICAO", "").
	// 	Exec()

	// fmt.Println("ID:", ds.ParamByName("ID").AsInt64())
	// //fmt.Println("Descrição", ds.ParamByName("OUT_DESCRICAO").AsString())
}

func TestDataSetMasterDetail(t *testing.T) {

	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnectionOracle(connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//ds1 := NewDataSet(db)
	//ds1.
	//	AddSql("select id, descricao from fab_processo").
	//	AddSql("where id = :id").
	//	SetInputParam("id", 41).
	//	Open()
	//
	//ds2 := NewDataSet(db)
	//ds2.AddSql("select id, codigo, descricao, id_processo from fab_operacao").
	//	AddMasterSource(ds1).
	//	AddDetailFields("id_processo").
	//	AddMasterFields("id").
	//	Open()
	//
	//fmt.Println("Processo:", ds1.FieldByName("descricao").AsString())
	//
	//for !ds2.Eof() {
	//	fmt.Println("Operações do processo:", ds2.FieldByName("descricao").AsString())
	//	ds2.Next()
	//}
}

func TestResultSetCLob(t *testing.T) {

	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnection(DialectType(ORACLE), connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//ds := NewDataSet(db)
	//
	//json := "{\n   \"basketHistory\":[\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":\"139\",\n                  \"groupName\":\"Este veículo foi rebocado na sequência de uma avaria?\",\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":\"Não\",\n                        \"contextId\":\"417\",\n                        \"contextValue\":\"Yes\"\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"1\",\n         \"basketcreationdate\":\"2023-10-10 15:57:23.0\",\n         \"dealernamelist\":\"EURO RIBEIRAO PRETO\"\n      },\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":\"139\",\n                  \"groupName\":\"Este veículo foi rebocado na sequência de uma avaria?\",\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":\"Não\",\n                        \"contextId\":\"417\",\n                        \"contextValue\":\"yes\"\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"2\",\n         \"basketcreationdate\":\"2023-10-09 21:29:16.0\",\n         \"dealernamelist\":\"EURO RIBEIRAO PRETO\"\n      },\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":\"139\",\n                  \"groupName\":\"Este veículo foi rebocado na sequência de uma avaria?\",\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":\"Não\",\n                        \"contextId\":\"417\",\n                        \"contextValue\":\"Yes\"\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"3\",\n         \"basketcreationdate\":\"2023-10-09 15:12:43.0\",\n         \"dealernamelist\":\"EURO RIBEIRAO PRETO\"\n      },\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":\"139\",\n                  \"groupName\":\"Este veículo foi rebocado na sequência de uma avaria?\",\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":\"Não\",\n                        \"contextId\":\"417\",\n                        \"contextValue\":\"Yes\"\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"4\",\n         \"basketcreationdate\":\"2023-10-09 14:17:31.0\",\n         \"dealernamelist\":\"EURO RIBEIRAO PRETO\"\n      },\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":\"139\",\n                  \"groupName\":\"Este veículo foi rebocado na sequência de uma avaria?\",\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":\"Não\",\n                        \"contextId\":\"417\",\n                        \"contextValue\":\"Yes\"\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"5\",\n         \"basketcreationdate\":\"2023-10-06 23:25:39.0\",\n         \"dealernamelist\":\"EURO RIBEIRAO PRETO\"\n      },\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":null,\n                  \"groupName\":null,\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":null,\n                        \"contextId\":null,\n                        \"contextValue\":null\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"6\",\n         \"basketcreationdate\":\"2023-09-19 21:46:49.0\",\n         \"dealernamelist\":\"EURO RIBEIRAO PRETO\"\n      },\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":\"139\",\n                  \"groupName\":\"Este veículo foi rebocado na sequência de uma avaria?\",\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":\"Não\",\n                        \"contextId\":\"417\",\n                        \"contextValue\":\"Yes\"\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"7\",\n         \"basketcreationdate\":\"2023-09-19 21:46:33.0\",\n         \"dealernamelist\":\"EURO RIBEIRAO PRETO\"\n      },\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":\"139\",\n                  \"groupName\":\"Este veículo foi rebocado na sequência de uma avaria?\",\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":\"Não\",\n                        \"contextId\":\"417\",\n                        \"contextValue\":\"Yes\"\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"8\",\n         \"basketcreationdate\":\"2023-09-15 19:46:48.0\",\n         \"dealernamelist\":null\n      },\n      {\n         \"functionLimits\":null,\n         \"docOrDiagnosisPresent\":\"1\",\n         \"recurrent\":\"0\",\n         \"numberOfInfotheques\":\"0\",\n         \"symptomId\":\"B019_001_001\",\n         \"rcCode\":\"2D\",\n         \"symptomCode\":\"B269\",\n         \"symptomLongLabel\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\",\n         \"eicCode\":null,\n         \"eicLabel\":null,\n         \"rcLabel\":null,\n         \"domainNodeInfo\":{\n            \"code\":\"B_019_001\",\n            \"label\":\"≤ 50 km/h (circuito urbano)\"\n         },\n         \"functionNodeInfo\":{\n            \"code\":\"B019_001_001\",\n            \"label\":\"Ruído/vibrações constantes em andamento a uma velocidade estabilizada: ≤ 50 km/h (circuito urbano)\"\n         },\n         \"symptomNodeInfo\":{\n            \"code\":\"B019\",\n            \"label\":\"A uma velocidade estabilizada\"\n         },\n         \"symptomContexts\":{\n            \"comments\":\"Test\",\n            \"contextGroupsArray\":[\n               {\n                  \"groupId\":\"139\",\n                  \"groupName\":\"Este veículo foi rebocado na sequência de uma avaria?\",\n                  \"comments\":null,\n                  \"adminContextArray\":[\n                     {\n                        \"contextLabel\":\"Não\",\n                        \"contextId\":\"417\",\n                        \"contextValue\":\"Yes\"\n                     }\n                  ]\n               }\n            ]\n         },\n         \"ecrCodes\":[\n            {\n               \"code\":\"1120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"120\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1225\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1720\",\n               \"label\":null\n            },\n            {\n               \"code\":\"1725\",\n               \"label\":null\n            },\n            {\n               \"code\":\"621\",\n               \"label\":null\n            },\n            {\n               \"code\":\"622\",\n               \"label\":null\n            },\n            {\n               \"code\":\"826\",\n               \"label\":null\n            },\n            {\n               \"code\":\"827\",\n               \"label\":null\n            }\n         ],\n         \"date\":\"2023-01-13 19:25:19.0\",\n         \"prestationNodeInfo\":{\n            \"code\":\"B\",\n            \"label\":\"Problemas com ruídos e vibração\"\n         },\n         \"basketno\":\"9\",\n         \"basketcreationdate\":\"2023-09-11 16:36:10.0\",\n         \"dealernamelist\":null\n      }\n   ],\n   \"vin\":\"VF1BM0B0H32443747\"\n}"
	////json := "teste"
	//_, err = ds.
	//	AddSql("INSERT INTO fab_mov_art_rena_ards (arquivo)").
	//	AddSql("values (:arquivo) returning id into :out_id").
	//	SetInputParamClob("arquivo", json).
	//	SetOutputParam("out_id", int64(0)).
	//	Exec()
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//id := ds.ParamByName("out_id").AsInt64()
	//
	//t.Log(id)
}

func TestDataSetToSInsertTransaction(t *testing.T) {

	t.Log("Sucesso.")

	//connectStr := "oracle://nbs_portal:new@100.0.65.225:1521/NBS1"
	//
	//db, err := NewConnection(DialectType(ORACLE), connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//tx, err := db.StartTransactionContext(context.Background())
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//ds := tx.NewDataSet()
	//
	//_, err = ds.
	//	AddSql("DELETE FROM USUARIO WHERE ID = 121").
	//	ExecContext(context.Background())
	//
	//if err != nil {
	//	tx.Rollback()
	//	fmt.Println("erro ao comitar.")
	//	t.Fatal(err)
	//}
	//
	//tx.Commit()
	//
	//fmt.Println("commitado com sucesso.")
}

func TestDataSetPostgres(t *testing.T) {

	t.Log("Sucesso.")

	connectStr := "postgres://postgres:100651xpto@localhost:5432/erp?sslmode=disable"

	db, err := NewConnection(DialectType(POSTGRESQL), connectStr)
	db.EnableLog()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	if err != nil {
		t.Fatal(err)
	}
	//
	//ds := db.NewDataSet()
	//
	//_, err = ds.
	//	AddSql("INSERT INTO PARAMETRO (ID, MODULO, CHAVE, DESCRICAO, TIPO, DADO, USUARIO, PERFIL, EMPRESA)").
	//	AddSql("VALUES (NEXTVAL('SEQ_PARAM_ID'),").
	//	AddSql("4, 'habilita_politica_senha','Habilita Política de Senha: 0 ou vazio = Não, 1 = Sim', 2, '0',0,0,0").
	//	AddSql(")").
	//	Exec()
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//for !ds.Eof() {
	//	fmt.Println(ds.FieldByName("ATIVO").AsBool())
	//
	//	ds.Next()
	//}
}

func TestDataSetParametroNull(t *testing.T) {

	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnectionOracle(connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//ds1 := NewDataSet(db)
	//ds1.
	//	AddSql("select id, descricao from fab_processo").
	//	AddSql("where (id = :id or :id is null)").
	//	SetInputParam("id", 41).
	//	Open()
	//
	//ds2 := NewDataSet(db)
	//ds2.AddSql("select id, codigo, descricao, id_processo from fab_operacao").
	//	AddMasterSource(ds1).
	//	AddDetailFields("id_processo").
	//	AddMasterFields("id").
	//	Open()
	//
	//fmt.Println("Processo:", ds1.FieldByName("descricao").AsString())
	//
	//for !ds2.Eof() {
	//	fmt.Println("Operações do processo:", ds2.FieldByName("descricao").AsString())
	//	ds2.Next()
	//}
}

func TestDataSetSelect(t *testing.T) {

	t.Log("Sucesso.")

	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := NewConnectionOracle(connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	ds.
		AddSql("select id, descricao from fab_processo").
		AddSql("where id = :id").
		SetInputParam("id", 19).
		Open()

	fmt.Println("Processo:", ds.FieldByName("descricao").AsString())
}

func TestDataSetExecProcedure(t *testing.T) {
	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnectionOracle(connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//var param1 int = 1
	//var param2 int = 10
	//
	//ds1 := NewDataSet(db)
	//_, err = ds1.
	//	AddSql("begin").
	//	AddSql(":result := pkg_crm_service.salvar_kit(:param1,:param2);").
	//	AddSql("end;").
	//	SetInputParam("param1", param1).
	//	SetInputParam("param2", param2).
	//	SetOutputParam("result", string("")).
	//	Exec()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println("sucesso: " + ds1.ParamByName("result").AsString())
}

func TestDataSetSqlWith(t *testing.T) {
	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnection(DialectType(ORACLE), connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//ds := NewDataSet(db)
	//
	//sql := `with
	//		  tb as (
	//		select
	//			'Plugin '|| ':modulo' || (select ' v.'|| VERSAO_MAJOR||'.'||VERSAO_MINOR from versao_modulo where cod_modulo = 594 ) as "versao"
	//		   ,to_char(sysdate,'rrrr-mm-dd')||'T'||to_char(sysdate,'hh24:mi:ss')||'.447Z' as "dataExtracaoDMS"
	//		   ,'NBS' as "dms"
	//		   ,o.numero_os ||'-'|| odv.chassi as "identificadorPassagem"
	//		from fab_mov_ei_TEST_mb_mgt a
	//		left join fab_ei_TEST_mb_mgt b on b.id = a.id_ei_TEST_mb_mgt
	//		left join os o on b.cod_empresa = o.cod_empresa
	//					  and b.numero_os = nvl(o.numero_os_fabrica,o.numero_os)
	//					  and o.status_os not in (5,6)
	//		left join os_dados_veiculos odv on o.cod_empresa = odv.cod_empresa
	//									   and o.numero_os = odv.numero_os
	//		where a.id_mov_mb_mgt = 1
	//		order by o.numero_os)
	//		select "versao" , "dataExtracaoDMS" , "dms" , "identificadorPassagem" from tb
	//		where rownum < 2`
	//
	//err = ds.
	//	AddSql(sql).
	//	Open()
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//fmt.Println("Processo:", ds.FieldByName("dataExtracaoDMS").AsString())
}

func TestDataSetExecFunction(t *testing.T) {

	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnection(DialectType(ORACLE), connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//dtini := time.Now()
	//dtfim := time.Now().AddDate(0, 0, 5)
	//
	//ds := NewDataSet(db)
	//_, err = ds.
	//	AddSql("begin").
	//	AddSql("  :DATAA := PKG_AGENDAMENTO_OFICINA.GET_FIM_CHIP(:DTINI,").
	//	AddSql("		   :DTFIM,").
	//	AddSql("		   1,").
	//	AddSql("		   :BOX,").
	//	AddSql("		   :MSG);").
	//	AddSql("end;").
	//	SetInputParam("DTINI", dtini).
	//	SetInputParam("DTFIM", dtfim).
	//	SetInputParam("BOX", "01001").
	//	SetOutputParam("MSG", string("")).
	//	SetOutputParam("DATAA", time.Now()).
	//	Exec()
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//fmt.Println(ds.ParamByName("MSG").AsString())
	//fmt.Println(ds.ParamByName("DATAA").AsString())
	//fmt.Println(ds.ParamByName("DATAA").AsDateTime())
}

func TestDataSetDateTime(t *testing.T) {
	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"
	//
	//db, err := NewConnectionOracle(connectStr)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//ds := NewDataSet(db)
	//err = ds.
	//	AddSql("select '06/12/2023 17:25:10' as vdata from dual").
	//	Open()
	//
	//fmt.Println(ds.findFieldByName("vdata").AsDateTime())
}

func TestSelectMySql(t *testing.T) {
	t.Log("Sucesso.")

	//connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	connectStr := "datasales:Da7a53735!C7M!@tcp(35.94.127.67:3306)/datasales?checkConnLiveness=false&maxAllowedPacket=0&parseTime=true"

	db, err := NewConnectionMySql(connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ds := NewDataSet(db)
	err = ds.
		AddSql("select id_bot_campanha_config, status, data_inicio_envio, data_fim_envio").
		AddSql("from bot_campanha_config where id_bot_campanha_config = :id_bot_campanha_config").
		SetInputParam("id_bot_campanha_config", 869).
		Open()

	fmt.Println(ds.FieldByName("data_inicio_envio").AsDateTime())
}
