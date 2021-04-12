package main

import "fmt"

func isPalindrome(s string) bool {
	n := len(s)
	i, j := 0, n-1
	for i <= j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	n := len(s)
	longest := 0
	longestPalin := ""
	for i := 0; i < n; i++ {
		for j := n; j >= i; j-- {
			if isPalindrome(s[i:j]) {
				if len(s[i:j]) > longest {
					longest = len(s[i:j])
					longestPalin = s[i:j]
					if i == 0 && j == n {
						break
					}
				}
			}
		}
	}
	return longestPalin
}

func main() {
	input := "babad"
	checkIsPalin(input)
	longestPalin(input)
	input = "babab"
	checkIsPalin(input)
	longestPalin(input)
	input = "malayalam"
	checkIsPalin(input)
	longestPalin(input)
	input = "a"
	checkIsPalin(input)
	longestPalin(input)
	input = "ba"
	checkIsPalin(input)
	longestPalin(input)
	input = "aa"
	checkIsPalin(input)
	longestPalin(input)
	input = "aaa"
	checkIsPalin(input)
	longestPalin(input)
}

func checkIsPalin(input string) {
	fmt.Printf("input = %v, output = %v\n", input, isPalindrome(input))
}

func longestPalin(input string) {
	fmt.Printf("input = %v, output = %v\n", input, longestPalindrome(input))
}
