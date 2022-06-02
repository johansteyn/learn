package datastructures

import (
	"fmt"
)

type Stack struct {
	list LinkedList
}

func NewStack() Stack {
	return Stack{NewList()}
}

func (s *Stack) Size() int {
	return s.list.size
}

func (s *Stack) Push(value int) {
	s.list.Add(value)
}

func (s *Stack) Pop() (int, error) {
	if s.list.Size() <= 0 {
		return 0, fmt.Errorf("empty stack")
	}
	value, err := s.list.Get(s.list.Size() - 1)
	if err != nil {
		return 0, err
	}
	err = s.list.Remove(s.list.Size() - 1)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if s.list.Size() <= 0 {
		return 0, fmt.Errorf("empty stack")
	}
	return s.list.Get(s.list.Size() - 1)
}

func (s *Stack) String() string {
	return s.list.String()
}
