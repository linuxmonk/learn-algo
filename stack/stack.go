package stack

import "errors"

type Stack struct {
	elems []interface{}
	size  int
	top   int
}

func New(sz int) *Stack {
	return &Stack{
		elems: make([]interface{}, sz, sz),
		size:  sz,
		top:   -1,
	}
}

func (s *Stack) Push(elem interface{}) error {
	if s == nil {
		panic("invalid stack")
	}
	if s.top == s.size-1 {
		return errors.New("stack is full")
	}
	s.top++
	s.elems[s.top] = elem
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s == nil {
		panic("invalid stack")
	}
	if s.top == -1 {
		return nil, errors.New("empty stack")
	}
	elem := s.elems[s.top]
	s.top--
	return elem, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s == nil {
		panic("invalid stack")
	}
	if s.top == -1 {
		return nil, errors.New("empty stack")
	}
	return s.elems[s.top], nil
}

func (s *Stack) Empty() bool {
	if s == nil {
		panic("invalid stack")
	}
	if s.top == -1 {
		return true
	}
	return false
}
