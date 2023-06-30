package godata

import (
	"database/sql"
	"time"
)

type Field struct {
	Name       string
	Caption    string
	DataType   *sql.ColumnType
	Value      variant
	DataMask   string
	BoolValue  bool
	TrueValue  string
	FalseValue string
	Visible    bool
	AcceptNull bool
	StrNull    string
	Order      int
	Index      int
}

type Fields []map[string]Field

func NewField(name string, dataType *sql.ColumnType) Field {
	field := Field{
		Name:       name,
		Caption:    name,
		DataType:   dataType,
		Value:      variant{},
		DataMask:   "",
		BoolValue:  false,
		TrueValue:  "",
		FalseValue: "",
		Visible:    true,
		AcceptNull: true,
		StrNull:    "null",
		Order:      1,
		Index:      0,
	}

	return field
}

func (field Field) AsValue() any {
	return field.Value.AsValue()
}

func (field Field) AsString() string {
	return field.Value.AsString()
}

func (field Field) AsInt() int {
	return field.Value.AsInt()
}

func (field Field) AsInt8() int8 {
	return field.Value.AsInt8()
}

func (field Field) AsInt16() int16 {
	return field.Value.AsInt16()
}

func (field Field) AsInt32() int32 {
	return field.Value.AsInt32()
}

func (field Field) AsInt64() int64 {
	return field.Value.AsInt64()
}

func (field Field) AsFloat() float32 {
	return field.Value.AsFloat()
}

func (field Field) AsFloat64() float64 {
	return field.Value.AsFloat64()
}

func (field Field) AsBool() bool {
	return field.Value.AsBool()
}

func (field Field) AsDateTime() time.Time {
	return field.Value.AsDateTime()
}

func (field Field) IsNull() bool {
	return field.Value.IsNull()
}

func (field Field) IsNotNull() bool {
	return field.Value.IsNull()
}
