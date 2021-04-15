package romantoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {

	tests := []struct {
		roman    string
		expected int
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
		v := RomanToInt(test.roman)
		assert.Equal(t, test.expected, v)
	}
}
