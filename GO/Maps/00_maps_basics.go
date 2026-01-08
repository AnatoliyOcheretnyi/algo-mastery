package main

import "fmt"

/*
Maps basics (super simple).

Tasks:
1) Create a map[string]int with two entries and print it.
2) Return the value for a key (assume it exists).
3) Increment the value for a key (create if missing).
4) Count frequency of words in a slice.
*/

func getValue(m map[string]int, key string) int {
	val, prs := m[key]
	if prs {
		return val
	}
	return 0
}

func incValue(m map[string]int, key string) {
	m[key] = m[key] + 1

}

func countWords(words []string) map[string]int {
	m := map[string]int{}

	for i := range words {

		m[words[i]] = m[words[i]] + 1
	}
	return m
}

func countWordsAnalog(words []string) map[string]int {
	m := map[string]int{}
	for _, w := range words {
		m[w]++
	}
	return m
}

func main() {
	// Task 1
	counts := map[string]int{"go": 1, "js": 2}
	fmt.Println("Task 1:", counts, "expected", "map[go:1 js:2]")

	// Task 2
	fmt.Println("Task 2:", getValue(counts, "go"), "expected", 1)

	// Task 3
	incValue(counts, "go")
	incValue(counts, "python")
	fmt.Println("Task 3:", counts, "expected", "map[go:2 js:2 python:1]")

	// Task 4
	fmt.Println("Task 4:", countWords([]string{"go", "js", "go"}), "expected", "map[go:2 js:1]")
}
