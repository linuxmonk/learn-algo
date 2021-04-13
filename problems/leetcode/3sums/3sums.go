package main

import "fmt"

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	out := threeSum(input)
	fmt.Println(out)
}

func threeSum(nums []int) [][]int {
	out := make([][]int, 0)
	if len(nums) == 0 {
		return out
	}
	if len(nums) < 3 {
		return out
	}
	processedIndex := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if i != j && i != k && j != k {
					if nums[i]+nums[j]+nums[k] == 0 {
						if processedIndex[i] && processedIndex[j] && processedIndex[k] {
							continue
						}
						sum := make([]int, 0)
						sum = append(sum, nums[i], nums[j], nums[k])
						out = append(out, sum)
						processedIndex[i] = true
						processedIndex[j] = true
						processedIndex[k] = true
					}
				}
			}
		}
	}
	return out
}
