package main

import (
	"fmt"
	"strconv"
)

/*
For loop basics.

Tasks:
1) Sum numbers from 1 to n (inclusive).
2) Count how many even numbers are in a slice.
3) Build a string with numbers 1..n separated by dashes.
*/

func sumToN(n int) int {
	i := 1
	result := 0

	for i <= n {
		result = result + i
		i++

	}
	return result
}

func countEvens(nums []int) int {
	result := 0

	for _, v := range nums {
		if v%2 == 0 {
			result ++
		}
	}
	return result
}

func joinWithDashes(n int) string {
	result := ""
	i := 1

	for i <= n {
		if i > 1 {
		result += "-"
	}
		result = result + strconv.Itoa(i)
		i++

	}
	return result
}

func main() {
	fmt.Println("Task 1:", sumToN(5), "expected", 15)
	fmt.Println("Task 1:", sumToN(1), "expected", 1)

	fmt.Println("Task 2:", countEvens([]int{1, 2, 3, 4, 6}), "expected", 3)
	fmt.Println("Task 2:", countEvens([]int{1, 3, 5}), "expected", 0)

	fmt.Println("Task 3:", joinWithDashes(1), "expected", "1")
	fmt.Println("Task 3:", joinWithDashes(4), "expected", "1-2-3-4")
}
