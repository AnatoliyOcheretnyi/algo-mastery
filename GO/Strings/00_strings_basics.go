package main

import "fmt"

/*
Strings basics (simple).

Tasks:
1) reverseString(s) -> reverse string.
2) countVowels(s) -> count vowels (a,e,i,o,u).
*/

func reverseString(s string) string {
	// TODO: implement
	return ""
}

func countVowels(s string) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println("Task 1:", reverseString("go"), "expected", "og")
	fmt.Println("Task 1:", reverseString("abc"), "expected", "cba")

	fmt.Println("Task 2:", countVowels("hello"), "expected", 2)
	fmt.Println("Task 2:", countVowels("sky"), "expected", 0)
}
