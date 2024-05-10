package parcial

import "fmt"

type ArraySet[T comparable] struct {
	elements []T
}

func NewArraySet[T comparable](elements ...T) *ArraySet[T] {
	set := &ArraySet[T]{make([]T, 0)}
	set.Add(elements...)
	return set
}

func (s *ArraySet[T]) Contains(element T) bool {
	for _, value := range s.elements {
		if element == value {
			return true
		}
	}
	return false
}

func (s *ArraySet[T]) Add(elements ...T) {
	for _, value := range elements {
		if !s.Contains(value) {
			s.elements = append(s.elements, value)
		}
	}
}

func (s *ArraySet[T]) Remove(element T) {
	i := 0
	for i < len(s.elements) {
		if s.elements[i] == element {
			break
		}
		i++
	}
	for i < len(s.elements)-1 {
		s.elements[i] = s.elements[i+1]
		i++
	}
	s.elements = s.elements[:len(s.elements)-1]
}

func (s *ArraySet[T]) Size() int {
	return len(s.elements)
}

func (s *ArraySet[T]) Values() []T {
	return s.elements
}

func (s *ArraySet[T]) String() string {
	str := "Set: {"
	for i, v := range s.elements {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%v", v)
	}
	str += "}"
	return str
}
