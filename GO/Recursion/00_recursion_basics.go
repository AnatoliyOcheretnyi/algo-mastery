package main

import "fmt"

/*
Recursion basics (simple).

Tasks:
1) factorial(n) -> n! (assume n >= 0)
2) sumToNRec(n) -> sum 1..n (assume n >= 0)
3) fib(n) -> n-th Fibonacci (0,1,1,2,3...) (assume n >= 0)
4) isPalindrome(s) -> true if string is palindrome (recursive)
*/

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func sumToNRec(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sumToNRec(n-1)
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-2) + fib(n-1)
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return isPalindrome(s[1 : len(s)-1])
}

func main() {
	fmt.Println("Task 1:", factorial(5), "expected", 120)
	fmt.Println("Task 1:", factorial(0), "expected", 1)

	fmt.Println("Task 2:", sumToNRec(5), "expected", 15)
	fmt.Println("Task 2:", sumToNRec(0), "expected", 0)

	fmt.Println("Task 3:", fib(0), "expected", 0)
	fmt.Println("Task 3:", fib(1), "expected", 1)
	fmt.Println("Task 3:", fib(5), "expected", 5)

	fmt.Println("Task 4:", isPalindrome("level"), "expected", true)
	fmt.Println("Task 4:", isPalindrome("hello"), "expected", false)
}
