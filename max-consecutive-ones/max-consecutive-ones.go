package main

import (
	"fmt"
)

func main() {
	nums, k := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2

	start, zeroCount, res := 0, 0, 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zeroCount++
		}
		for ; zeroCount > k; start++ {
			if nums[start] == 0 {
				zeroCount--
			}
		}
		if (end - start + 1) > res {
			res = end - start + 1
		}
	}
	fmt.Println(res)
}
