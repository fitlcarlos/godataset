package godataset

import (
	"fmt"
	"reflect"
	"strings"
)

type Macros struct {
	Owner *DataSet
	List  []*Macro
}

type Macro struct {
	Name  string
	Value *Variant
}

func NewMacros() *Macros {
	value := &Macros{
		List: []*Macro{},
	}
	return value
}

func (m *Macros) FindMacroByName(macroName string) *Macro {
	for i := 0; i < len(m.List); i++ {
		if strings.ToUpper(m.List[i].Name) == strings.ToUpper(macroName) {
			return m.List[i]
		}
	}
	return nil
}

func (m *Macros) MacroByName(macroName string) *Macro {
	var macro *Macro

	for i := 0; i < len(m.List); i++ {
		if strings.ToUpper(m.List[i].Name) == strings.ToUpper(macroName) {
			macro = m.List[i]
		}
	}

	if macro == nil {
		macro = &Macro{}
		fmt.Println("Macro " + macroName + " doesn't exists")
	}

	return macro
}

func (m *Macros) SetMacro(macroName string, macroValue any) *Macros {
	macro := m.FindMacroByName(macroName)

	if macro != nil {
		macro.Value.Value = macroValue
	} else {
		macro = &Macro{
			Name:  macroName,
			Value: &Variant{Value: macroValue},
		}
		m.List = append(m.List, macro)
	}

	return m
}

func (m *Macros) Count() int {
	return len(m.List)
}

func (m *Macros) Clear() *Macros {
	m.Owner = nil
	ClearSlice(m.List)
	m.List = nil
	return m
}

func (macro *Macro) AsValue() *Variant {
	if macro.Value == nil {
		return &Variant{}
	}

	if macro.Value.Value == nil {
		return macro.Value
	}

	if IsPointer(macro.Value.Value) {
		a := reflect.ValueOf(macro.Value.Value).Elem().Interface()
		return &Variant{Value: a}
	} else {
		return macro.Value
	}
}

func (macro *Macro) AsString() string {
	return macro.AsValue().AsString()
}
