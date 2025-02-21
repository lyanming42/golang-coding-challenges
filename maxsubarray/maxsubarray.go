package main

import (
	"fmt"
	"slices"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}

type queue []int

func (q queue) Enqueue(v int) queue {
	return append(q, v)
}

func (q queue) Dequeue() (queue, int) {
	l := len(q)
	if l == 0 {
		return q, 0
	}
	return q[1:], q[0]
}

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printPerson(pers Person) {
	fmt.Println(pers.age)
	fmt.Println(pers.name)
	fmt.Println(pers.job)
	fmt.Println(pers.salary)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	const PI = 3.14

	s := make(stack, 0)

	s = s.Push(10)
	s = s.Push(11)

	s, p := s.Pop()
	fmt.Println(p)

	q := make(queue, 0)

	q = q.Enqueue(10)
	q = q.Enqueue(10)

	q, item := q.Dequeue()
	fmt.Println(item)

	var nums []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
	}
	max := slices.Max(nums)
	fmt.Println(max)

	var pers1 Person

	pers1.age = 10
	pers1.name = "abcd"
	pers1.job = "hkr"
	pers1.salary = 12

	fmt.Println(pers1.age)
	fmt.Println(pers1.name)
	fmt.Println(pers1.job)
	fmt.Println(pers1.salary)

	printPerson((pers1))

	var a = map[string]string{"Ford": "Ford", "Team": "Chelsea"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Println(a["Ford"])
	fmt.Println(b)

	fmt.Println(factorial_recursion(4))

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	copiedNumbers := make([]int, 2, 2)
	copy(copiedNumbers, myslice)

	fmt.Println(myslice)
	fmt.Println(copiedNumbers)

	fmt.Println(&nums)
}
