package godataset

import (
	"reflect"
	"time"
)

type ParamType int

const (
	IN    ParamType = 0
	OUT   ParamType = 1
	INOUT ParamType = 2
)

type Value interface{ *Variant | Variant }
type Param struct {
	Owner     *Params
	Name      string
	Value     *Variant
	ParamType ParamType
	DataType  reflect.Type
	Values    []*Variant
}

func NewParam(paramName string, paramType ParamType) *Param {
	return &Param{
		Name:      paramName,
		ParamType: paramType,
		Values:    []*Variant{},
	}
}
func (param *Param) AsValue() *Variant {
	if param.Value == nil {
		return &Variant{}
	}

	if param.Value.Value == nil {
		return param.Value
	}

	if IsPointer(param.Value.Value) {
		a := reflect.ValueOf(param.Value.Value).Elem().Interface()
		return &Variant{Value: a}
	} else {
		return param.Value
	}
}

func (param *Param) AsString() string {
	return param.AsValue().AsString()
}

func (param *Param) AsInt() int {
	return param.AsValue().AsInt()
}

func (param *Param) AsInt8() int8 {
	return param.AsValue().AsInt8()
}

func (param *Param) AsInt16() int16 {
	return param.AsValue().AsInt16()
}

func (param *Param) AsInt32() int32 {
	return param.AsValue().AsInt32()
}

func (param *Param) AsInt64() int64 {
	return param.AsValue().AsInt64()
}

func (param *Param) AsFloat() float32 {
	return param.AsValue().AsFloat()
}

func (param *Param) AsFloat64() float64 {
	return param.AsValue().AsFloat64()
}

func (param *Param) AsBool() bool {
	return param.AsValue().AsBool()
}

func (param *Param) AsDateTime() time.Time {
	return param.AsValue().AsDateTime()
}
