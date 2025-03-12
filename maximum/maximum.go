package main

import (
	"fmt"
)

func main() {
	freqs := map[byte]int{10: 22, 12: 12, 13: 24, 22: 10}
	max := 0

	for _, item := range freqs {
		if max < item {
			max = item
		}
	}

	fmt.Println(max)
}
