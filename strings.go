package godata

import (
	"strings"
)

type Strings struct {
	Items []string
}

func New() *Strings {
	return &Strings{}
}

func (s *Strings) Append(value string) *Strings {
	s.Items = append(s.Items, value)
	return s
}

func (s *Strings) Clear() *Strings {
	ClearSlice(s.Items)
	s.Items = nil
	return s
}

func (s *Strings) Add(value string) *Strings {
	s.Items = append(s.Items, value)
	return s
}

func (s *Strings) Replace(index int, value string) *Strings {
	s.Items[index] = value
	return s
}

func (s *Strings) Count() int {
	return len(s.Items)
}

func (s *Strings) Text() string {
	return strings.Join(s.Items, " \n")
}

func (s *Strings) AddStrings(value *Strings) {
	for i := 0; i < value.Count(); i++ {
		s.Add(value.Items[i])
	}
}
