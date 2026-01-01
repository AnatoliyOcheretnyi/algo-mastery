package main

import "fmt"

/*
Basics: data types and containers in Go.

Types to know:
- string: text
- int: whole numbers
- bool: true/false
- []T: slice (dynamic length)
- [N]T: array (fixed length)
- map[K]V: dictionary (key/value)

Notes:
- Arrays have fixed length (part of the type).
- Slices are views over arrays and can grow.
*/

func main() {
	// Task 1: create variables of basic types and print them.
	// Example: name := "Ada"
	// TODO: create: name (string), age (int), active (bool)
	var name = "Anatolii"
	var age = 26
	var active = true

	name = "Tolik"

	// Task 2: array vs slice.
	// TODO: create an array of 3 ints, e.g. [3]int{1, 2, 3}

	var arr = [3]int{1, 2, 3}
	// TODO: create a slice of ints, e.g. []int{1, 2, 3}

	var slice = []int{1, 2, 3}

	// Task 3: append to a slice (arrays cannot append).
	// TODO: append one value to the slice

	slice = append(slice, slice...)

	// Task 4: map basics.
	// TODO: create a map[string]int and put two key/value pairs
	// Example: counts["go"] = 1

	counts := make(map[string]int)

	counts["go"] = 1
	counts["js"] = 2

	// Print everything so you can see it.
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("active:", active)
	fmt.Println("array:", arr)
	fmt.Println("slice:", slice)
	fmt.Println("map:", counts)
}
