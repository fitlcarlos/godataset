package godata

type Row struct {
	List map[string]Variant
}

func NewRow() Row {
	row := Row{
		List: map[string]Variant{},
	}
	return row
}

func (r *Row) Clear() *Row {
	r.List = nil
	return r
}
