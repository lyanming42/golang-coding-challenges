package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
	}
	max := slices.Max(nums)
	fmt.Println(max)
}
