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
7) setNilToZero(p **int): if *p is nil, allocate and set to 0.
8) appendViaPointer(nums *[]int, v int): append value to slice via pointer.
9) resetUser(u *User): set Name="" and Age=0 via pointer.
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

func setNilToZero(p **int) {
	// TODO: implement
}

func appendViaPointer(nums *[]int, v int) {
	// TODO: implement
}

func resetUser(u *User) {
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

	// Task 7
	var p *int
	setNilToZero(&p)
	fmt.Println("Task 7:", *p, "expected", 0)

	// Task 8
	list := []int{1, 2}
	appendViaPointer(&list, 3)
	fmt.Println("Task 8:", list, "expected", []int{1, 2, 3})

	// Task 9
	u2 := User{Name: "Bob", Age: 40}
	resetUser(&u2)
	fmt.Println("Task 9:", u2, "expected", "{ 0}")
}
