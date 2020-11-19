package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string
	details := make(map[string]string)

	// Set up the input reader:
	scanner := bufio.NewScanner(os.Stdin)
	
	// Prompt the user to enter a name:
	fmt.Println("Please enter a name...")

	// Get the user's input:
	scanner.Scan()
	name = scanner.Text()

	// Check that a valid name was entered:
	if name == "" {
		fmt.Println("ERROR: Invalid name - please enter a non-blank value!")
		return
	}

	// Prompt the user to enter an address:
	fmt.Println("Please enter an address...")

	// Get the user's input:
	scanner.Scan()
	address = scanner.Text()
	
	// Check that a valid name was entered:
	if address == "" {
		fmt.Println("ERROR: Invalid address - please enter a non-blank value!")
		return
	}

	// Construct the map:
	details["name"] = name
	details["address"] = address

	// Serialize the output to JSON:
	outputJson, err := json.Marshal(details)

	// Check that the serialization succeeded:
	if err != nil {
		fmt.Println("ERROR: Serialization to JSON failed!")
		return
	}

	outputString := string(outputJson)

	// Print the output:
	fmt.Println(outputString)
}