package godata

type Row struct {
	List map[string]variant
}

func NewRow() Row {
	row := Row{
		List: make(map[string]variant),
	}
	return row
}
