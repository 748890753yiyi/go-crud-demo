package main

type Set struct {
	m map[interface{}]struct{}
}
var Exists = struct{}{}
func New(items ...interface{}) *Set  {
	s := &Set{}
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

func (s *Set) Add(items ...interface{}) error  {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear()  {
	s.m = make(map[interface{}]struct{})
}

func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() {
		return false
	}
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set) IsSubset(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}