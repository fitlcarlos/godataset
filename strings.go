package godata

import (
	"strings"
)

type Strings struct {
	Items []string
}

func (s *Strings) Append(value string) *Strings {
	s.Items = append(s.Items, value)
	return s
}
func (s *Strings) Clear() *Strings {
	s.Items = nil
	return s
}
func (s *Strings) Add(value string) *Strings {
	s.Items = append(s.Items, value)
	return s
}
func (s *Strings) Count() int {
	return len(s.Items)
}
func (s *Strings) Text() string {
	return strings.Join(s.Items, " \n")
}
