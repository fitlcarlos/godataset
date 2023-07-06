package godata

import "strings"

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
