package godata

import (
	"strings"
)

type Fields struct {
	Owner *DataSet
	List  map[string]*Field
}

func NewFields() Fields {
	fields := Fields{
		List: make(map[string]*Field),
	}
	return fields
}
func (fd *Fields) FieldByName(fieldName string) *Field {
	field, ok := fd.List[strings.ToUpper(fieldName)]
	if ok {
		return field
	} else {
		return nil
	}
}

func (fd *Fields) Add(fieldName string) *Field {
	field := fd.FieldByName(fieldName)

	if field != nil {
		return field
	} else {
		field = NewField(fieldName)
		field.Owner = fd
		fd.List[strings.ToUpper(fieldName)] = field
		return field
	}
}
