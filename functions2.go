package main
import "fmt"

func add(a int, b int) int { return a + b }
func subtract(a, b int) int { return a - b }

func execute(a, b int, operation func(int, int) int) int { // shorthand for duplicate types
	return operation(a, b)
}

func main() {
	i, j := 4, 2

	added := execute(i, j, add)
	fmt.Printf("Added %d + %d == %d\n", i, j, added)

	subtracted := execute(i, j, subtract)
	fmt.Printf("Subtracted %d - %d == %d\n", i, j, subtracted)
}
