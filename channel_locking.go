package main
import "fmt"

func fibonacci(i uint) uint {
	if i <= 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	fmt.Println(fibonacci(41))
	fmt.Println(fibonacci(41))
}
