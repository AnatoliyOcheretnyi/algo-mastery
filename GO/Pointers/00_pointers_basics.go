package main

import "fmt"

/*
Pointers basics (simple).

Tasks:
1) zeroVal(v int): try to set v to 0 (value copy).
2) zeroPtr(p *int): set the value at pointer to 0.
3) swap(a, b *int): swap two integers via pointers.
4) incrementPtr(p *int): add 1 to the value via pointer.
5) updateName(u *User, name string): update struct field via pointer.
6) setFirst(nums []int, v int): set first element (note: slices are references).
*/

type User struct {
	Name string
	Age  int
}

func zeroVal(v int) {
	// TODO: implement
}

func zeroPtr(p *int) {
	// TODO: implement
}

func swap(a, b *int) {
	// TODO: implement
}

func incrementPtr(p *int) {
	// TODO: implement
}

func updateName(u *User, name string) {
	// TODO: implement
}

func setFirst(nums []int, v int) {
	// TODO: implement
}

func main() {
	// Task 1/2
	x := 5
	zeroVal(x)
	fmt.Println("Task 1:", x, "expected", 5)
	zeroPtr(&x)
	fmt.Println("Task 2:", x, "expected", 0)

	// Task 3
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println("Task 3:", a, b, "expected", 2, 1)

	// Task 4
	incrementPtr(&a)
	fmt.Println("Task 4:", a, "expected", 3)

	// Task 5
	u := User{Name: "Ada", Age: 20}
	updateName(&u, "Grace")
	fmt.Println("Task 5:", u.Name, "expected", "Grace")

	// Task 6
	nums := []int{9, 8, 7}
	setFirst(nums, 1)
	fmt.Println("Task 6:", nums, "expected", []int{1, 8, 7})
}
