package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Response struct {
	Pattern  string `json:"pattern"`
	Response string `json:"response"`
}

var QUIT_WORDS = []string{"exit", "quit", "bye", "goodbye", "goodnight", "good morning", "night", "morning"}
var RESPONSES []Response
var FALLBACK = "WHAT"

func loadRules(filename string) ([]Response, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var responses []Response
	err = json.Unmarshal(byteValue, &responses)
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func main() {
	rules, err := loadRules("rules.json")
	if err != nil {
		fmt.Printf("Error loading rules: %v\n", err)
		os.Exit(1)
	}
	RESPONSES = rules

	fmt.Println("How do you do. Please tell me your problem.")
	var input string
	var response string

	reader := bufio.NewReader(os.Stdin)

	for !slices.Contains(QUIT_WORDS, strings.ToLower(input)) {
		fmt.Print("> ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println(input)

		for _, r := range RESPONSES {
			if matched, _ := regexp.MatchString(r.Pattern, input); matched {
				response = r.Response
				break
			}
		}

		if slices.Contains(QUIT_WORDS, strings.ToLower(input)) {
			response = "Goodbye. I hope our conversation was helpful."
			break
		}

		if response == "" {
			response = FALLBACK
		}

		fmt.Println(response)
	}

}
