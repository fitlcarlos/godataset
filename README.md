# GoDataset

GODataset - Biblioteca Go para facilitar a interação com bancos de dados e manipulação de dados via dataset, inspirada em conceitos de datasets de outras linguagens, mas com a simplicidade e performance do Go.

## Sumário
- [Descrição](#descrição)
- [Instalação](#instalação)
- [Exemplo de Uso](#exemplo-de-uso)
- [Principais Métodos](#principais-métodos)
- [Contribuição](#contribuição)
- [Referências](#referências)
- [Licença](#licença)
- [Autor](#autor)

## Descrição
O GoDataset abstrai operações comuns de manipulação de dados, permitindo trabalhar com registros, campos e navegação de forma eficiente. Ideal para aplicações que precisam de flexibilidade e produtividade ao lidar com diferentes bancos de dados.

## Instalação

Requer Go 1.18 ou superior.

```bash
go get github.com/seuusuario/godataset
```

## Exemplo de Uso
```go
connectStr := "oracle://erp:pass123@DESKTOP-DEV:1521/xe"

db, err := NewConnection(DialectType(ORACLE), connectStr)
if err != nil {
    log.Fatal(err)
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
    log.Fatal(err)
}

fmt.Println(ds.Count())
ds.First()
for !ds.Eof() {
    fmt.Println(ds.FieldByName("NAME").AsString())
    ds.Next()
}
```

## Principais Métodos

- `NewDataSet(db *Conn) *DataSet`
- `NewDataSetTx(tx *Transaction) *DataSet`
- `(*DataSet) AddSql(sql string) *DataSet`
- `(*DataSet) SetInputParam(paramName string, paramValue any) *DataSet`
- `(*DataSet) Open() error`
- `(*DataSet) Exec() (sql.Result, error)`
- `(*DataSet) FieldByName(fieldName string) *Field`
- `(*DataSet) First()`, `Next()`, `Previous()`, `Last()`
- `(*DataSet) Eof() bool`, `Bof() bool`, `IsEmpty() bool`, `Count() int`
- `(*DataSet) ToStruct(model any) error`

Veja a documentação completa no código-fonte para todos os métodos disponíveis.

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

1. Fork este repositório
2. Crie uma branch para sua feature (`git checkout -b feature/nome-feature`)
3. Commit suas alterações (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nome-feature`)
5. Abra um Pull Request

## Referências
[![Go Reference](https://pkg.go.dev/badge/github.com/fitlcarlos/godataset.svg)](https://pkg.go.dev/github.com/fitlcarlos/godataset)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Licença
Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Autor
Carlos Fitl

---
Sinta-se à vontade para sugerir melhorias ou reportar problemas!