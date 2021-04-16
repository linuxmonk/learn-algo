package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {

	tests := []struct {
		input    int
		expected int
	}{
		{
			3,
			6,
		},
		{
			4,
			24,
		},
		{
			0,
			0,
		},
		{
			-1,
			-1,
		},
		{
			1,
			1,
		},
	}
	for _, test := range tests {
		fact := Factorial(test.input)
		assert.Equal(t, test.expected, fact)
	}
}
