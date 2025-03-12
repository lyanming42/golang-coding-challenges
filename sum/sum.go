package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("Sum= %v\n", sum)

	array1, array2 := [3]int{1, 2, 3}, [3]int{1, 2, 4}

	if array1 == array2 {
		fmt.Println("True")
	}
}
