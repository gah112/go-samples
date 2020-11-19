package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	// Declare a string to hold the user's input:
	var input string

	// Set up the input reader:
	scanner := bufio.NewScanner(os.Stdin)

	// Instantiate the animal objects:
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

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
		
		// Warn the user if the input does not contain 2 fields:
		if len(strs) < 2 {
			fmt.Println("WARNING: You must enter both an animal and an action separated by a space. Continuing...")
			continue
		}

		if(strs[0] == "cow") {
			if(strs[1] == "eat") {
				cow.Eat()
			} else if (strs[1] == "move") {
				cow.Move()
			} else if (strs[1] == "speak") {
				cow.Speak()
			} else {
				fmt.Println("WARNING: Invalid action - you must provide an action of 'eat', 'move', or 'speak'. Continuing...")
				continue
			}
		} else if(strs[0] == "bird") {
			if(strs[1] == "eat") {
				bird.Eat()
			} else if (strs[1] == "move") {
				bird.Move()
			} else if (strs[1] == "speak") {
				bird.Speak()
			} else {
				fmt.Println("WARNING: Invalid action - you must provide an action of 'eat', 'move', or 'speak'. Continuing...")
				continue
			}
		} else if(strs[0] == "snake") {
			if(strs[1] == "eat") {
				snake.Eat()
			} else if (strs[1] == "move") {
				snake.Move()
			} else if (strs[1] == "speak") {
				snake.Speak()
			} else {
				fmt.Println("WARNING: Invalid action - you must provide an action of 'eat', 'move', or 'speak'. Continuing...")
				continue
			}
		} else {
			fmt.Println("WARNING: Invalid animal - you must provide an animal of 'cow', 'bird' or 'snake'. Continuing...")
			continue
		}
	}
}