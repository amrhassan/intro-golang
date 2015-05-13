package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Tick!")
		time.Sleep(1000)
		fmt.Println("Tock!")
		time.Sleep(1000)
	}
}
