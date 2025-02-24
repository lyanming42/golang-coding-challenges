package main

import (
	"fmt"
	"math"
)

// window sliding algorithm demonstration

func main() {
	nums := []int{1, 4, 4}
	left := 0
	windowSum := 0
	target := 4

	res := math.MaxInt64
	for right := 0; right < len(nums); right++ {
		windowSum += nums[right]
		for ; windowSum >= target; left++ {
			windowSum -= nums[left]
			if res > (right - left + 1) {
				res = right - left + 1
			}
		}
	}
	fmt.Println(res)
}
