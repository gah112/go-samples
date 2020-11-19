package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the Animal interface with Eat(), Move(), and Speak() methods:
type Animal interface {
	Name() string
	Eat()
	Move()
	Speak()
}

// Define the Cow type with a name parameter:
type Cow struct {
	name string
}
func (cow Cow) Name() string { return cow.name }
func (cow Cow) Eat() { fmt.Println("grass") }
func (cow Cow) Move() { fmt.Println("walk") }
func (cow Cow) Speak() { fmt.Println("moo") }

// Define the Bird type with a name parameter:
type Bird struct {
	name string
}
func (bird Bird) Name() string { return bird.name }
func (bird Bird) Eat() { fmt.Println("worms") }
func (bird Bird) Move() { fmt.Println("fly") }
func (bird Bird) Speak() { fmt.Println("peep") }

// Define the Snake type with a name parameter:
type Snake struct {
	name string
}
func (snake Snake) Name() string { return snake.name }
func (snake Snake) Eat() { fmt.Println("mice") }
func (snake Snake) Move() { fmt.Println("slither") }
func (snake Snake) Speak() { fmt.Println("hsss") }

func main() {
	// Declare a string to hold the user's input:
	var input string

	// Declare a list of Animals to hold the user's items:
	var animals []Animal

	// Set up the input reader:
	scanner := bufio.NewScanner(os.Stdin)

	for (true) {
		// Prompt the user to enter a string:
		fmt.Println("Please enter an animal and an action...")

		// Get the user input:
		scanner.Scan()
		input = strings.ToLower(scanner.Text())

		// Break if the escape character has been entered:
		if input == "x" {
			break
		}

		// Split the string to separate the animal and action:
		strs := strings.Split(input," ")
		
		// Warn the user if the input does not contain 3 fields:
		if len(strs) < 3 {
			fmt.Println("WARNING: You must enter 3 arguments separated by a space. Continuing...")
			continue
		}

		// Check the 1st argument:
		if(strs[0] == "newanimal") {
			if(strs[1] == "") {
				fmt.Println("WARNING: Invalid name - you cannot provide an empty name when creating a new animal. Continuing...")
				continue
			} else {
				// Create the animal:
				if(strs[2] == "cow") {
					animal := Cow{strs[1]}
					animals = append(animals,animal)
					fmt.Println("Created it!")
					continue
				} else if(strs[2] == "bird") {
					animal := Bird{strs[1]}
					animals = append(animals,animal)
					fmt.Println("Created it!")
					continue
				} else if(strs[2] == "snake") {
					animal := Snake{strs[1]}
					animals = append(animals,animal)
					fmt.Println("Created it!")
					continue
				} else {
					fmt.Println("WARNING: Invalid animal - you must provide an animal of 'cow', 'bird', or 'snake'. Continuing...")
					continue
				}
			}
		} else if(strs[0] == "query") {
			if(strs[1] == "") {
				fmt.Println("WARNING: Invalid name - you cannot provide an empty name when querying for an animal. Continuing...")
				continue
			} else if(strs[2] != "eat" && strs[2] != "move" && strs[2] != "speak") {
				fmt.Println("WARNING: Invalid action - you must provide an action of 'eat', 'move', or 'speak' when querying for an animal. Continuing...")
				continue
			} else {
				// Declare a boolean to hold the result of the user's query:
				found := false

				// Query for the animal by name:
				for _,animal := range(animals) {
					if(strs[1] == animal.Name()) {
						found = true
						if(strs[2] == "eat") {
							animal.Eat()
							continue
						} else if(strs[2] == "move") {
							animal.Move()
							continue
						} else if(strs[2] == "speak") {
							animal.Speak()
							continue
						}
					}
				}
				if(!found) {
					fmt.Println("INFO: Animal not found - have you created an animal with the queried name? Continuing...")
					continue
				}
			}
		} else {
			fmt.Println("WARNING: Invalid entry - you must enter a 1st parameter of 'newanimal' or 'query'. Continuing...")
			continue
		}
	}
}