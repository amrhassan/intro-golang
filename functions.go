package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(3, 1)
	fmt.Println("Initial sum", sum)

	// Closures! and first-class function values
	incrementSum := func() { sum += 1 }

	incrementSum()
	incrementSum()
	fmt.Println("Incremented sum", sum)
}
