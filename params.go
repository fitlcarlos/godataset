package godata

type Params struct {
	value map[string]Param
}

func NewParams() *Params {
	value := &Params{
		value: make(map[string]Param),
	}
	return value
}
func (p Params) ParamByName(paramName string) Param {
	return p.value[paramName]
}
