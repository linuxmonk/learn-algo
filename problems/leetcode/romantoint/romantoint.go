package romantoint

import (
	"fmt"
	"strings"
)

var romanValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var subtractions = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func RomanToInt(roman string) int {
	upper := strings.ToUpper(roman)
	n := len(upper)
	ret := 0
	for i := 0; i < n; i++ {
		v := romanValues[upper[i]]
		if i+1 < n && v < romanValues[upper[i+1]] {
			key := fmt.Sprintf("%c%c", upper[i], upper[i+1])
			ret = ret + subtractions[key]
			i++
		} else {
			ret = ret + v
		}
	}
	return ret
}
