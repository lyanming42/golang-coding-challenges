package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var i, j int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{i, j}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, -11, 5, 4}, 9))
}
