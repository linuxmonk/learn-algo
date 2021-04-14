package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {

	n := len(s)
	num := 0
	foundNums := false
	sign := 1
	i := 0
	done := false
	// eat white spaces
	for s[i] == ' ' {
		i++
	}
	for ; i < n; i++ {
		if done {
			break
		}
		switch {
		case s[i] == '-':
			if !foundNums {
				sign = -1
			}
		case s[i] >= '0' && s[i] <= '9':
			foundNums = true
			num = (10 * num) + int(s[i]-'0')
		default:
			done = true
			break
		}
	}
	ret := num * sign
	if ret <= math.MinInt32 {
		return math.MinInt32
	}
	if ret >= math.MaxInt32 {
		return math.MaxInt32
	}
	return num * sign
}

func main() {
	input := "32"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "00456"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "-42"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "      -42"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "4193 with words"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "words and 987"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = " words with two numbers -9345 and 987"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "-91283472332"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
}
