package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string;
	
	// Set up the input reader:
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt entry from the user:
	fmt.Println("Please enter a string:")

	// Get the user input:
	scanner.Scan()
	input = scanner.Text()

	// Print the result of pattern matching:
	if strings.Index(strings.ToLower(input),"i") == 0 && strings.LastIndex(strings.ToLower(input),"n") == len(input) - 1 && strings.Contains(strings.ToLower(input),"a") {
				fmt.Println("Found!")
				return
	} else {
		fmt.Println("Not Found!")
		return
	}
}