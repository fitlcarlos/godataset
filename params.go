package godataset

import (
	"fmt"
	goOra "github.com/sijms/go-ora/v2"
	"reflect"
	"strings"
	"time"
)

type Params struct {
	Owner     *DataSet
	BatchSize int
	List      []*Param
}

func NewParams() *Params {
	value := &Params{
		List: []*Param{},
	}
	return value
}

func (p *Params) FindParamByName(paramName string) *Param {
	var param *Param
	for i := 0; i < len(p.List); i++ {
		if strings.ToUpper(p.List[i].Name) == strings.ToUpper(paramName) {
			param = p.List[i]
		}
	}

	return param
}

func (p *Params) ParamByName(paramName string) *Param {
	if param := p.FindParamByName(paramName); param != nil {
		return param
	}

	fmt.Println("Parameter " + paramName + " doesn't exists")
	return &Param{}
}

func (p *Params) Add(paramName string) *Param {
	param := p.FindParamByName(paramName)

	if param != nil {
		return param
	} else {
		param = &Param{
			Name: paramName,
		}
		p.List = append(p.List, param)
		return param
	}
}

func (p *Params) SetParam(paramName string, paramValue any, paramType ParamType) *Params {

	param := p.FindParamByName(paramName)

	switch paramType {
	case IN, INOUT:
		if param != nil {
			param.Value.Value = paramValue
		} else {
			param = &Param{
				Name:      paramName,
				Value:     &Variant{Value: paramValue},
				ParamType: paramType,
			}
			p.List = append(p.List, param)
		}
	case OUT:
		if param != nil {
			param.Value.Value = paramValue
		} else {
			switch paramValue.(type) {
			case int, int8, int16, int32, int64:
				value := int64(0)
				param = &Param{
					Name:      paramName,
					Value:     &Variant{Value: &value},
					ParamType: OUT,
				}
			case float32:
				value := float32(0)
				param = &Param{
					Name:      paramName,
					Value:     &Variant{Value: &value},
					ParamType: OUT,
				}
			case float64:
				value := float64(0)
				param = &Param{
					Name:      paramName,
					Value:     &Variant{Value: &value},
					ParamType: OUT,
				}
			case string:
				value := generateString()
				param = &Param{
					Name:      paramName,
					Value:     &Variant{Value: &value},
					ParamType: OUT,
				}
			case bool:
				value := false
				param = &Param{
					Name:      paramName,
					Value:     &Variant{Value: &value},
					ParamType: OUT,
				}
			case time.Time:
				value := time.Time{}
				param = &Param{
					Name:      paramName,
					Value:     &Variant{Value: &value},
					ParamType: OUT,
				}
			default:
				value := float64(0)
				param = &Param{
					Name:      paramName,
					Value:     &Variant{Value: &value},
					ParamType: OUT,
				}
			}

			p.List = append(p.List, param)
		}
	}

	return p
}

func (p *Params) SetParamClob(paramName string, paramValue string, paramType ParamType) *Params {
	if p.Owner.Connection.Dialect == ORACLE {
		p.SetParam(paramName, goOra.Clob{String: paramValue, Valid: StrNotEmpty(paramValue)}, paramType)
	} else if p.Owner.Connection.Dialect == POSTGRESQL {
		p.SetParam(paramName, []byte(paramValue), paramType)
	} else {
		p.SetParam(paramName, paramValue, paramType)
	}
	return p
}

func (p *Params) SetParamBlob(paramName string, paramValue []byte, paramType ParamType) *Params {
	if p.Owner.Connection.Dialect == ORACLE {
		p.SetParam(paramName, goOra.Blob{Data: paramValue, Valid: len(paramValue) > 0}, paramType)
	} else if p.Owner.Connection.Dialect == POSTGRESQL {
		p.SetParam(paramName, paramValue, paramType)
	} else {
		p.SetParam(paramName, paramValue, paramType)
	}
	return p
}

func (p *Params) SetInputParam(paramName string, paramValue any) *Params {
	return p.SetParam(paramName, paramValue, IN)
}

func (p *Params) SetInputParamClob(paramName string, paramValue string) *Params {
	return p.SetParamClob(paramName, paramValue, IN)
}

func (p *Params) SetInputParamBlob(paramName string, paramValue []byte) *Params {
	return p.SetParamBlob(paramName, paramValue, IN)
}

func (p *Params) SetInOutputParam(paramName string, paramValue any) *Params {
	return p.SetParam(paramName, paramValue, INOUT)
}

func (p *Params) SetInOutputParamClob(paramName string, paramValue string) *Params {
	return p.SetParamClob(paramName, paramValue, INOUT)
}

func (p *Params) SetInOutputParamBlob(paramName string, paramValue []byte) *Params {
	return p.SetParamBlob(paramName, paramValue, INOUT)
}

func (p *Params) SetOutputParam(paramName string, paramValue any) *Params {
	return p.SetParam(paramName, paramValue, OUT)
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

func (p *Params) Count() int {
	return len(p.List)
}

func (p *Params) Clear() *Params {
	p.Owner = nil

	for i := 0; i < len(p.List); i++ {
		p.List[i].Owner = nil
	}

	ClearSlice(p.List)
	p.List = nil
	return p
}
