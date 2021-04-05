package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPushPop(t *testing.T) {

	s := New(8)
	s.Push(1)
	s.Push(2)
	v, err := s.Pop()
	assert.Equal(t, err, nil)
	i := v.(int)
	assert.Equal(t, i, 2)

	v, err = s.Pop()
	assert.Equal(t, err, nil)
	i = v.(int)
	assert.Equal(t, i, 1)

	empty := s.Empty()
	assert.Equal(t, true, empty)
}

func TestStackEmptyError(t *testing.T) {
	s := New(4)
	v, err := s.Pop()
	assert.Equal(t, v, nil)
	assert.NotEqual(t, err, nil)
}

func TestStackPeek(t *testing.T) {
	s := New(4)
	s.Push(1)
	s.Push(2)
	v, err := s.Peek()
	assert.Equal(t, err, nil)
	i := v.(int)
	assert.Equal(t, i, 2)
	v, err = s.Peek()
	assert.Equal(t, err, nil)
	i = v.(int)
	assert.Equal(t, i, 2)

	empty := s.Empty()
	assert.Equal(t, false, empty)
}

func TestStackParens(t *testing.T) {
	s := New(32)
	expr := "()(){()}"
	for _, e := range expr {
		if e == '(' || e == '{' {
			s.Push(e)
		}
		if e == ')' || e == '}' {
			s.Pop()
		}
	}
	empty := s.Empty()
	assert.Equal(t, true, empty)
}
