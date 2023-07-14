package godata

import (
	"fmt"
	"reflect"
	"strings"
)

type Params struct {
	List map[string]Param
}

func NewParams() *Params {
	value := &Params{
		List: make(map[string]Param),
	}
	return value
}
func (p *Params) ParamByName(paramName string) Param {
	return p.List[paramName]
}

func (p *Params) Add(paramName string) Param {
	param := p.ParamByName(paramName)

	if &param != nil {
		return param
	} else {
		param = Param{
			Name: paramName,
		}
		p.List[strings.ToUpper(paramName)] = param
		return param
	}
}

func (p *Params) SetOutputParam(paramName string, paramType any) *Params {
	switch paramType.(type) {
	case int, int8, int16, int32, int64:
		dataType := int64(0)
		p.List[paramName] = Param{Value: Variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case float32:
		dataType := float32(0)
		p.List[paramName] = Param{Value: Variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case float64:
		dataType := float64(0)
		p.List[paramName] = Param{Value: Variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	case string:
		dataType := generateString()
		p.List[paramName] = Param{Value: Variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	default:
		dataType := float64(0)
		p.List[paramName] = Param{Value: Variant{Value: &dataType}, DataType: reflect.TypeOf(dataType), ParamType: OUT}
	}
	return p
}

func (p *Params) PrintParam() {
	for key, value := range p.List {
		fmt.Println("Colum:", key, "Value:", value.AsValue(), "Type:", reflect.TypeOf(value.AsValue()))
	}
}
