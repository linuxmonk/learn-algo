package linter

import (
	"errors"
	"fmt"

	"github.com/gusaki/algorithms/stack"
)

type Linter struct {
	s *stack.Stack
}

func New() *Linter {
	l := &Linter{
		s: stack.New(1024),
	}
	return l
}

func (l *Linter) Lint(expr string) error {

	for _, c := range expr {
		if c == '(' || c == '{' || c == '[' {
			l.s.Push(c)
			continue
		}
		if c == ')' || c == '}' || c == ']' {
			v, err := l.s.Pop()
			if err != nil {
				return err
			}
			paren, _ := v.(rune)
			switch paren {
			case '(':
				if c != ')' {
					return fmt.Errorf("Unbalanced parenthesis. Expected ) found %c", c)
				}
			case '{':
				if c != '}' {
					return fmt.Errorf("Unbalanced braces. Expected } found %c", c)
				}
			case '[':
				if c != ']' {
					return fmt.Errorf("Unbalanced square brackets. Expected ] found %c", c)
				}
			}
		}
	}
	if !l.s.Empty() {
		return errors.New("Missing closing brackets/parenthesis")
	}
	return nil
}
