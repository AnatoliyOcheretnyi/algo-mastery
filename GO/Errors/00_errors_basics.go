package main

import (
	"errors"
	"fmt"
)

/*
Errors basics (simple).

Tasks:
1) divide(a, b) -> return result and error if b == 0.
2) parsePositive(s) -> parse int, error if <= 0.
*/

func divide(a, b int) (int, error) {
	// TODO: implement
	return 0, nil
}

func parsePositive(s string) (int, error) {
	// TODO: implement
	return 0, nil
}

func main() {
	fmt.Println("Task 1:", divide(10, 2), "expected", "5 <nil>")
	fmt.Println("Task 1:", divide(10, 0), "expected", "0 error")

	fmt.Println("Task 2:", parsePositive("7"), "expected", "7 <nil>")
	fmt.Println("Task 2:", parsePositive("-3"), "expected", "0 error")
	fmt.Println("Task 2:", parsePositive("abc"), "expected", "0 error")

	_ = errors.New
}
