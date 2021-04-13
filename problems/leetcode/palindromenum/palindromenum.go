package main

import "fmt"

func revInt(x int) int {
	rev := 0
	for x > 0 {
		rem := x % 10
		x = x / 10
		rev = (10 * rev) + rem
	}
	return rev
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x == revInt(x) {
		return true
	}
	return false
}
func main() {
	in := 121
	fmt.Println("Is palindrome number: ", isPalindrome(in))
	in = 1231
	fmt.Println("Is palindrome number: ", isPalindrome(in))
}
