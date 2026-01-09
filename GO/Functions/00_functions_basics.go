package main

import "fmt"

/*
Functions practice (basic).

Topics:
- Functions
- Multiple return values
- Variadic functions
- Closures
*/

// Task 1 (functions): return the square of a number.
func square(n int) int {
	return n * n
}

// Task 2 (multiple returns): return sum and product of two numbers.
func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b

}

// Task 3 (variadic): return the sum of all numbers.
func sumAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum = sum + n
	}
	return sum
}

// Task 4 (closures): return a function that adds x to its input.
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	fmt.Println("Task 1:", square(5), "expected", 25)

	sum, prod := sumAndProduct(3, 4)
	fmt.Println("Task 2:", sum, prod, "expected", 7, 12)

	fmt.Println("Task 3:", sumAll(1, 2, 3), "expected", 6)
	fmt.Println("Task 3:", sumAll(), "expected", 0)

	add10 := makeAdder(10)
	fmt.Println("Task 4:", add10(5), "expected", 15)
}
