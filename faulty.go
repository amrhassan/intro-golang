package main

import (
	"fmt"
)

func getMyText() (string, error) {
	// This could have been read from disk or a network socket
	return "Your text", nil
}

func main() {
	myText, err := getMyText()
	if err != nil {
		myText = "Sorry, I couldn't get your text. May I interest you in a joke?"
	}

	fmt.Println(myText)
}
