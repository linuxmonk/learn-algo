package inttoroman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {

	tests := []struct {
		expected string
		num      int
	}{
		{
			"I",
			1,
		},
		{
			"V",
			5,
		},
		{
			"III",
			3,
		},
		{
			"IV",
			4,
		},
		{
			"IX",
			9,
		},
		{
			"LVIII",
			58,
		},
		{
			"MCMXCIV",
			1994,
		},
	}
	for _, test := range tests {
		v := IntToRoman(test.num)
		assert.Equal(t, test.expected, v)
	}
}
