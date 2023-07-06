package godata

import (
	"database/sql"
	"strings"
	"time"
)

type Field struct {
	Owner    *Fields
	Name     string
	Caption  string
	DataType *sql.ColumnType
	//Value      variant
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

func NewField(name string) *Field {
	field := &Field{
		Name:    name,
		Caption: name,
		//Value:      variant{},
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
	return field.getVariant().AsValue()
}

func (field Field) AsString() string {
	return field.getVariant().AsString()
}

func (field Field) AsInt() int {
	return field.getVariant().AsInt()
}

func (field Field) AsInt8() int8 {
	return field.getVariant().AsInt8()
}

func (field Field) AsInt16() int16 {
	return field.getVariant().AsInt16()
}

func (field Field) AsInt32() int32 {
	return field.getVariant().AsInt32()
}

func (field Field) AsInt64() int64 {
	return field.getVariant().AsInt64()
}

func (field Field) AsFloat() float32 {
	return field.getVariant().AsFloat()
}

func (field Field) AsFloat64() float64 {
	return field.getVariant().AsFloat64()
}

func (field Field) AsBool() bool {
	return field.getVariant().AsBool()
}

func (field Field) AsDateTime() time.Time {
	return field.getVariant().AsDateTime()
}

func (field Field) IsNull() bool {
	return field.getVariant().IsNull()
}

func (field Field) IsNotNull() bool {
	return field.getVariant().IsNotNull()
}

func (field Field) getVariant() variant {
	index := field.Owner.Owner.Index
	return field.Owner.Owner.Rows[index].List[strings.ToUpper(field.Name)]
}
