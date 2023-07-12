package godata

type Row struct {
	List map[string]Variant
}

func NewRow() Row {
	row := Row{
		List: make(map[string]Variant),
	}
	return row
}
