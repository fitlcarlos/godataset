package godata

import (
	"database/sql"
	"time"
)

type Field struct {
	Name       string
	Caption    string
	DataType *sql.ColumnType
	Value    variant
	DataMask string
	ValueTrue  string
	ValueFalse string
	Visible    bool
	Order      int
	Index      int
}

type Fields []map[string]Field

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
