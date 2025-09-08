package datastructures

import (
	"fmt"
	"iter"
	"maps"
)

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func NewSetFrom[T comparable](values ...T) *Set[T] {
	s := NewSet[T]()
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s *Set[T]) Iter() iter.Seq[T] {
	return maps.Keys(s.m)
}

func (s *Set[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, found := s.m[value]
	return found
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) List() []T {
	keys := make([]T, 0, len(s.m))
	for key := range s.m {
		keys = append(keys, key)
	}
	return keys
}

func (s *Set[T]) String() string {
	return fmt.Sprintf("%v", s.List())
}

func (s *Set[T]) Filter(p func(T) bool) *Set[T] {
	r := NewSet[T]()
	for e := range s.m {
		if p(e) {
			r.Add(e)
		}
	}
	return r
}

// We can only map to a set of the same type - not to a different type
func (s *Set[T]) Map(fn func(T) T) *Set[T] {
	r := NewSet[T]()
	for e := range s.m {
		r.Add(fn(e))
	}
	return r
}

// Functions can return different types
func MapSet[T1 comparable, T2 comparable](s *Set[T1], fn func(T1) T2) *Set[T2] {
	r := NewSet[T2]()
	for e := range s.m {
		r.Add(fn(e))
	}
	return r
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	r := NewSet[T]()
	for e := range s.m {
		r.Add(e)
	}
	for e := range s2.m {
		r.Add(e)
	}
	return r
}

func (s *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	r := NewSet[T]()
	ref := s
	other := s2
	if len(s2.m) < len(s.m) {
		ref, other = s2, s
	}

	for e := range ref.m {
		if _, ok := other.m[e]; ok {
			r.Add(e)
		}
	}

	return r
}

func (s *Set[T]) Subtract(s2 *Set[T]) *Set[T] {
	r := NewSet[T]()
	for e := range s.m {
		if _, ok := s2.m[e]; !ok {
			r.Add(e)
		}
	}
	return r
}
