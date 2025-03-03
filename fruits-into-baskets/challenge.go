package main

import (
	"fmt"
)

func main() {
	fruits := []int{1, 2, 3, 2, 2}
	fruitsCount := map[int]int{}
	res := 0

	start := 0
	for end := 0; end < len(fruits); end++ {
		fruitsCount[fruits[end]]++

		for ; start <= end && len(fruitsCount) > 2; start++ {
			fruitsCount[fruits[start]]--

			if fruitsCount[fruits[start]] == 0 {
				delete(fruitsCount, fruits[start])
			}
		}
		if res < (end - start + 1) {
			res = end - start + 1
		}
	}

	fmt.Println(res)
}
