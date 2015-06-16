package main
import "fmt"

func compute(computation func() int) <- chan int {
	outputChannel := make(chan int)
	go func() {
		outputChannel <- computation()
	}()
	return outputChannel
}

func fibonacci(i int) int { if i <= 1 { return 1 } else {return fibonacci(i-1) + fibonacci(i-2) }}

func main() {
	computation := func() int {return fibonacci(41)}
	firstResult := compute(computation)
	secondResult := compute(computation)
	// Now I'm gonna block and wait for the two results
	fmt.Println(<- firstResult)
	fmt.Println(<- secondResult)
}