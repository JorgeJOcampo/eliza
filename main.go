package main

import "fmt"
import "slices"

var QUIT_WORDS = []string{"exit", "quit", "bye", "goodbye", "goodnight", "good morning", "night", "morning"}

func main() {
	fmt.Println("Hello! This is Eliza!\n")
	var input string
	for !slices.Contains(QUIT_WORDS, input) {
		fmt.Scan(&input)
	}
}
