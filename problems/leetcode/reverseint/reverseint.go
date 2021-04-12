package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
	}
	ret := 0
	x = x * sign
	for x >= 10 {
		rem := x % 10
		x = x / 10
		ret = (10 * ret) + (10 * rem)
	}
	ret = (ret + x) * sign
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret
}

func main() {
	input := 1234
	fmt.Printf("input = %v, output = %v\n", input, reverse(input))

	input = -1234
	fmt.Printf("input = %v, output = %v\n", input, reverse(input))

	input = 1350
	fmt.Printf("input = %v, output = %v\n", input, reverse(input))

	input = 10
	fmt.Printf("input = %v, output = %v\n", input, reverse(input))

	input = math.MinInt32
	fmt.Printf("input = %v, output = %v\n", input, reverse(input))

	input = math.MaxInt32
	fmt.Printf("input = %v, output = %v\n", input, reverse(input))

}
