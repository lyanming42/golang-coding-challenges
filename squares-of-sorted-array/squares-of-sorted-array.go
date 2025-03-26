package main

func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}

	result := make([]int, 0)

	firstPositiveIdx := len(nums)

	for idx, num := range nums {
		if num >= 0 {
			firstPositiveIdx = idx
			break
		}
	}

	left := firstPositiveIdx - 1
	right := firstPositiveIdx

	for right < len(nums) && left >= 0 {
		if nums[right] < Abs(nums[left]) {
			result = append(result, nums[right]*nums[right])
			right++

		} else {
			result = append(result, nums[left]*nums[left])
			left--
		}
	}
	for ; left >= 0; left-- {
		result = append(result, nums[left]*nums[left])
	}

	for ; right < len(nums); right++ {
		result = append(result, nums[right]*nums[right])
	}

}
