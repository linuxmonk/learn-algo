package linter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinterExpression(t *testing.T) {
	l := New()
	err := l.Lint("r = (3 / (2 - y)) + (x * x)")
	assert.Equal(t, nil, err)
}

func TestLinterErrorExpression(t *testing.T) {
	l := New()
	err := l.Lint("r = 3 / (2 - y)) + (x * x)")
	assert.NotEqual(t, err, nil)
}

func TestLinterCodeOK(t *testing.T) {
	code := `func TestGoFn(s string) {
		for _, v := range(s) {
			fmt.Printf("Rune: %v\n", v)
		}
	}`
	l := New()
	err := l.Lint(code)
	assert.Equal(t, nil, err)
}

func TestLinterBadCode(t *testing.T) {
	code := `func TestGoFn(s string {
		for _, v := range(s) {
			fmt.Printf("Rune: %v\n", v)
		}
	}`
	l := New()
	err := l.Lint(code)
	assert.NotEqual(t, nil, err)
}
