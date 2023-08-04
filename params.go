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

func (p *Params) SetInputParam(paramName string, paramValue any) *Params {
	p.List[paramName] = Param{Value: Variant{Value: paramValue}, ParamType: IN}
	return p
}

func (p *Params) SetOutputParam(paramName string, paramValue any) *Params {
	switch paramValue.(type) {
	case int, int8, int16, int32, int64:
		value := int64(0)
		p.List[paramName] = Param{Value: Variant{Value: &value}, ParamType: OUT}
	case float32:
		value := float32(0)
		p.List[paramName] = Param{Value: Variant{Value: &value}, ParamType: OUT}
	case float64:
		value := float64(0)
		p.List[paramName] = Param{Value: Variant{Value: &value}, ParamType: OUT}
	case string:
		value := generateString()
		p.List[paramName] = Param{Value: Variant{Value: &value}, ParamType: OUT}
	default:
		value := float64(0)
		p.List[paramName] = Param{Value: Variant{Value: &value}, ParamType: OUT}
	}
	return p
}

func (p *Params) SetOutputParamSlice(params ...ParamOut) *Params {
	for i := 0; i < len(params); i++ {
		p.SetOutputParam(params[i].Name, params[i].Dest)
	}
	return p
}

func (p *Params) PrintParam() {
	for key, value := range p.List {
		fmt.Println("Colum:", key, "Value:", value.AsValue(), "Type:", reflect.TypeOf(value.AsValue()))
	}
}
