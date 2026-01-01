package main

import (
	"fmt"
	"time"
)

// Task 1 (string): join first and last name with a space.
func fullName(first, last string) string {
	// TODO: implement
	return ""
}

// Task 2 (int, []int): sum all numbers in nums.
func sumInts(nums []int) int {
	// TODO: implement
	return 0
}

// Task 3 (bool, []bool): return true if all values are true.
func allTrue(values []bool) bool {
	// TODO: implement
	return false
}

// Task 4 (time.Time): return absolute difference in full days.
func daysBetween(a, b time.Time) int {
	// TODO: implement
	return 0
}

// Task 5 (map[string]int): count occurrences of each word.
func countOccurrences(words []string) map[string]int {
	// TODO: implement
	return map[string]int{}
}

func main() {
	fmt.Println("Task 1:", fullName("Ada", "Lovelace"), "expected", "Ada Lovelace")

	fmt.Println("Task 2:", sumInts([]int{1, 2, 3, 4}), "expected", 10)

	fmt.Println("Task 3:", allTrue([]bool{true, true, false}), "expected", false)
	fmt.Println("Task 3:", allTrue([]bool{true, true, true}), "expected", true)

	a, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00Z")
	b, _ := time.Parse(time.RFC3339, "2024-01-04T02:00:00Z")
	fmt.Println("Task 4:", daysBetween(a, b), "expected", 3)

	fmt.Println("Task 5:", countOccurrences([]string{"go", "js", "go", "go"}), "expected", "map[go:3 js:1]")
}
