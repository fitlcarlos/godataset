package godata

import (
	"reflect"
	"time"
)

type Input int

const (
	IN    Input = 0
	OUT   Input = 1
	INOUT Input = 2
)

type Value interface{ *variant | variant }
type Param struct {
	Value    variant
	Input    Input
	DataType reflect.Type
}

type Params map[string]Param

func (param Param) AsValue() variant {
	if IsPointer(param.Value.Value) {
		a := reflect.ValueOf(param.Value.Value).Elem().Interface()
		return variant{Value: a}
	} else {
		return param.Value
	}
}

func (param Param) AsString() string {
	return param.Value.AsString()
}

func (param Param) AsInt() int {
	return param.Value.AsInt()
}

func (param Param) AsInt8() int8 {
	return param.Value.AsInt8()
}

func (param Param) AsInt16() int16 {
	return param.Value.AsInt16()
}

func (param Param) AsInt32() int32 {
	return param.Value.AsInt32()
}

func (param Param) AsInt64() int64 {
	return param.Value.AsInt64()
}

func (param Param) AsFloat() float32 {
	return param.Value.AsFloat()
}

func (param Param) AsFloat64() float64 {
	return param.Value.AsFloat64()
}

func (param Param) AsBool() bool {
	return param.Value.AsBool()
}

func (param Param) AsDateTime() time.Time {
	return param.Value.AsDateTime()
}
