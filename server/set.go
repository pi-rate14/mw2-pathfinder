package main

type set struct {
	m map[string]Node
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]Node)
	return s
}

func (s *set) Add(value Node) {
	s.m[value.Source] = value
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}
