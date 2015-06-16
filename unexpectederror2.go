package main

import ("fmt"; "regexp")

func extractLongestNumber(text string) string {
	extractorRegexp := regexp.MustCompile("([0-9]+)")
	matches := extractorRegexp.FindStringSubmatch(text)
	if len(matches) > 1 { return matches[1] } else { return "" }
}

func main() {
	myText := "Sonmi-451"
	fmt.Println(extractLongestNumber(myText))
}