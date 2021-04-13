package main

import (
	"fmt"
	"sort"
)

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
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		threeSumII(nums, i, &out)
	}
	return out
}

func threeSumII(nums []int, i int, out *[][]int) {
	lo := i + 1
	hi := len(nums) - 1
	for lo < hi {
		threeSum := nums[lo] + nums[i] + nums[hi]
		if threeSum < 0 {
			lo++
			for {
				if lo < hi && nums[lo] == nums[lo+1] {
					lo++
					continue
				}
				break
			}
			continue
		}
		if threeSum > 0 {
			hi--
			for {
				if hi > lo && nums[hi] == nums[hi-1] {
					hi--
					continue
				}
				break
			}
		}
		if threeSum == 0 {
			*out = append(*out, []int{nums[lo], nums[i], nums[hi]})
			lo++
			for i == lo {
				lo++
			}
			hi--
		}
	}
}
