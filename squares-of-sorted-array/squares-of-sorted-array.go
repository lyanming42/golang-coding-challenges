package main

import "fmt"

func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}

	result := make([]int, len(nums))

	left := 0
	right := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		rightSq := nums[right] * nums[right]
		leftSq := nums[left] * nums[left]
		if rightSq >= leftSq {
			result[i] = rightSq
			right--
		} else {
			result[i] = leftSq
			left++
		}
	}
	fmt.Println(result)
}
