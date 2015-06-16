package main

import ("fmt"; "regexp")

func extractLongestNumber(text string) string {
	extractorRegexp, err := regexp.Compile("([0-9]+)")
	if err != nil {
		panic(err) // This will crash with a stack trace and the value of err
	}

	matches := extractorRegexp.FindStringSubmatch(text)
	if len(matches) > 1 { return matches[1] } else { return "" }
}

func main() {
	myText := "Sonmi-451"
	fmt.Println(extractLongestNumber(myText))
}
