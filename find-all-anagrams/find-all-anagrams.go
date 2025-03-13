package main

import "fmt"

func main() {
	s1, s2 := "ab", "eidbaooo"
	start, count := 0, [26]int{}
	res := make([]int, 0)

	for idx := range s1 {
		count[s1[idx]-97]++
	}

	for end := range s2 {
		count[s2[end]-97]--

		if (end - start + 1) > len(s1) {
			count[s2[start]-97]++
			start++
		}

		if count == [26]int{} {
			res = append(res, start)
		}
	}

	fmt.Println(res)
}
