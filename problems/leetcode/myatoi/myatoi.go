package main

import (
	"fmt"
	"math"
	"strings"
)

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isSign(c byte) bool {
	if c == '-' || c == '+' {
		return true
	}
	return false
}

func isNum(word string) bool {

	var dot bool

	if len(word) == 0 {
		return false
	}
	i := 0
	if isSign(word[0]) {
		i = 1
	}
	// could be a number without a +
	for ; i < len(word); i++ {
		if isDigit(word[i]) {
			continue
		}
		if word[i] == '.' && !dot {
			dot = true
			continue
		}
		return false
	}
	return true
}

func myAtoi(s string) int {

	var num, sign int

	s = strings.TrimSpace(s)
	sign = 1
	slist := strings.Split(s, " ")

	for _, word := range slist {
		if word == "" {
			continue
		}
		if !isNum(word) {
			return 0
		}
		i := 0
		if isSign(word[0]) {
			if word[0] == '-' {
				sign = -1
			}
			i = 1
		}
		for ; i < len(word); i++ {
			if word[i] == '.' {
				break
			}
			num = (10 * num) + int(word[i]-'0')
		}
		break
	}

	ret := num * sign
	if ret <= math.MinInt32 {
		return math.MinInt32
	}
	if ret >= math.MaxInt32 {
		return math.MaxInt32
	}
	return ret
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
	input = "+-12"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "-+12"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "+12"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
	input = "3.145632"
	fmt.Printf("input = %v, output = %v\n", input, myAtoi(input))
}
