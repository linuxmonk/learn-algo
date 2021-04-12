package main

import (
	"container/ring"
	"fmt"
)

func findTheWinner(n int, k int) int {
	r := ring.New(n)
	for i := 0; i < n; i++ {
		r.Value = i + 1
		r = r.Next()
	}
	for r.Len() != 1 {
		r.Do(func(p interface{}) {
			fmt.Printf("%d ", p.(int))
		})
		fmt.Println()
		for i := 1; i < k-1; i++ {
			r = r.Next()
		}
		r.Unlink(1)
	}
	return r.Value.(int)
}

func main() {
	winner := findTheWinner(5, 2)
	fmt.Println("n = 5, k = 2, winner = ", winner)
	// winner = findTheWinner(6, 5)
	// fmt.Println("n = 6, k = 5, winner = ", winner)
}
