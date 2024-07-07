package godata

type Row struct {
	List map[string]*Variant
}

func NewRow() Row {
	row := Row{
		List: map[string]*Variant{},
	}
	return row
}
