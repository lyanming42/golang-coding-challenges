package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	cur := nums[0] - 1

	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != cur {
			cur = nums[i]
			nums[cnt] = cur
			cnt++
		}
	}
	fmt.Println(nums, cnt)
}
