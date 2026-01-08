package main

import "fmt"

/*
Slices basics (super simple).

Tasks:
1) Create a slice with 3 integers and print it.
2) Append one element to the slice.
3) Return the length of a slice.
4) Return the last element of a slice (assume non-empty).
5) Sum all elements in a slice.
*/

func lengthOfSlice(nums []int) int {
	return len(nums)
}

func lastElement(nums []int) int {

	length := len(nums)
	return nums[length - 1]
}

func sumSlice(nums []int) int {
	sum := 0
	for i:= range nums {
		sum = sum + nums[i]
	}
	return sum
}

func main() {
	// Task 1
	nums := []int{1, 2, 3}

	// Task 2
	nums = append(nums, 4)

	// Task 3
	fmt.Println("Task 3:", lengthOfSlice(nums), "expected", 4)

	// Task 4
	fmt.Println("Task 4:", lastElement(nums), "expected", 4)

	// Task 5
	fmt.Println("Task 5:", sumSlice(nums), "expected", 10)
}
