package main

import "fmt"
import "slices"
import "regexp"

type Response struct {
	Pattern string
	Response string
}

var QUIT_WORDS = []string{"exit", "quit", "bye", "goodbye", "goodnight", "good morning", "night", "morning"}
var RESPONSES = []Response{
	{Pattern: "Hello", Response: "Hello! This is Eliza!\n"},
}

func main() {
	fmt.Println("Hello! This is Eliza!\n")
	var input string
	for !slices.Contains(QUIT_WORDS, input) {
		fmt.Scan(&input)
		for _, response := range RESPONSES {
			if matched, _ := regexp.MatchString(response.Pattern, input); matched {
				fmt.Println(response.Response)
				break
			}
		}
	}
}
