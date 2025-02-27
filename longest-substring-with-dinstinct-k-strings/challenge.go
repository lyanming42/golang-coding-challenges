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
	var s string = "abcdeffg"
	var chmap map[int]int = map[int]int{0: 0}

	for i := int('a'); i <= int('z'); i++ {
		chmap[i] = 0
	}

	fmt.Println(getUniqueChs(chmap))

	left, right := 0, 0

	for ; right < len(s); right++ {
		fmt.Println(s[right])
		left++
	}
}
