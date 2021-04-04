package main

import "fmt"

func main() {
	fmt.Println("abcabcbb", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbbb", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstring("pwwkew"))
	fmt.Println(" ", lengthOfLongestSubstring(" "))
	fmt.Println("aab", lengthOfLongestSubstring("aab"))
	fmt.Println("dvdf", lengthOfLongestSubstring("dvdf"))
}

type Counter struct {
	Count int
	Pos   int
}

func lengthOfLongestSubstring(s string) int {

	var c byte
	m := make(map[byte]*Counter)
	i, currlen, maxlen := 0, 0, 0

	for {
		if i > len(s)-1 {
			if currlen > maxlen {
				maxlen = currlen
			}
			break
		}
		c = s[i]
		_, ok := m[c]
		if !ok {
			counter := &Counter{
				Count: 1,
				Pos:   i,
			}
			m[c] = counter
			i++
			currlen++
			continue
		}
		oldpos := m[c].Pos
		i = oldpos + 1
		if currlen > maxlen {
			maxlen = currlen
		}
		m = make(map[byte]*Counter)
		currlen = 0
	}
	return maxlen
}
