package godata

import (
	"fmt"
	"strings"
)

type Fields struct {
	Owner *DataSet
	List  []*Field
}

func NewFields() *Fields {
	fields := &Fields{
		List: []*Field{},
	}

	return fields
}

func (fd *Fields) FindFieldByName(fieldName string) *Field {
	for i := 0; i < len(fd.List); i++ {
		if strings.EqualFold(fd.List[i].Name, fieldName) {
			return fd.List[i]
		}
	}
	return nil
}

func (fd *Fields) FieldByName(fieldName string) *Field {
	for i := 0; i < len(fd.List); i++ {
		if strings.EqualFold(fd.List[i].Name, fieldName) {
			return fd.List[i]
		}
	}

		fmt.Println("Field " + fieldName + " doesn't exists")
	return &Field{Owner: fd}
	}

func (fd *Fields) Add(fieldName string) (field *Field) {
	field = NewField(fieldName)
	field.Owner = fd

	ok := fd.FindFieldByName(fieldName) != nil
	if ok {
		field.Name = fmt.Sprintf("%s_%d", fieldName, fd.countRepeated(fieldName))
}

	fd.List = append(fd.List, field)
	return
}

func (fd *Fields) countRepeated(fieldName string) (counter int) {
	for i := 0; i < len(fd.List); i++ {
		if strings.EqualFold(fd.List[i].originalName, fieldName) {
			counter++
		}
	}
	return
}

func (fd *Fields) Clear() *Fields {
	fd.Owner = nil

	for i := 0; i < len(fd.List); i++ {
		fd.List[i].Owner = nil
	}

	ClearSlice(fd.List)
	fd.List = nil
	return fd
}

func (fd *Fields) Count() int {
	return len(fd.List)
}
