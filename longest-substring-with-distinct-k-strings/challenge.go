package main

import (
	"fmt"
)

func main() {
	var s string = "aabbbcc"
	var chmap map[int]int = make(map[int]int)

	k := 2
	left, right, res := 0, 0, 0

	for ; right < len(s); right++ {
		chmap[int(s[right])]++
		for ; len(chmap) > k && left <= right; left++ {
			chmap[int(s[left])]--
			if chmap[int(s[left])] == 0 {
				delete(chmap, int(s[left]))
			}
		}
		if len(chmap) == k {
			if res < (right - left + 1) {
				res = right - left + 1
			}
		}
	}
	fmt.Println(res)
}
