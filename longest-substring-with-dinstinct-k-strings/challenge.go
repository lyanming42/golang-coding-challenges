package main

import (
	"fmt"
)

func getUniqueChs(chmap map[int]int) int {
	cnt := 0
	for i := int('a'); i <= int('z'); i++ {
		if chmap[i] > 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	var s string = "aabbcc"
	var chmap map[int]int = map[int]int{0: 0}

	for i := int('a'); i <= int('z'); i++ {
		chmap[i] = 0
	}

	k := 1
	left, right, res := 0, 0, 0

	for ; right < len(s); right++ {
		chmap[int(s[right])] = chmap[int(s[right])] + 1
		if getUniqueChs(chmap) == k {
			if res < (right - left + 1) {
				res = right - left + 1
			}
		}
		for ; getUniqueChs(chmap) > k && left <= right; left++ {
			chmap[int(s[left])] = chmap[int(s[left])] - 1
		}
		if getUniqueChs(chmap) == k {
			if res < (right - left + 1) {
				res = right - left + 1
			}
		}
	}
	fmt.Println(res)
}
