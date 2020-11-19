package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var file string
	var persons []Person

	// Prompt the user to enter a file name:
	fmt.Println("Enter the name of the text file...")

	// Get the user's input:
	fmt.Scan(&file)

	// Try to read the file:
	content, err := ioutil.ReadFile(file)

	// Handle file errors:
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
	}

	// Convert the file bytes into a string:
	contentString := string(content)

	// Split the string lines:
	names := strings.Split(contentString, "\n")

	// Iteratively add each string as a person:
	for _,name := range names {
		words := strings.Split(name," ")
		persons = append(persons, Person{words[0],strings.Join(words[1:]," ")})
	}

	// Iteratively read from the list of persons:
	for _,p := range persons {
		fmt.Printf("FIRST NAME: %s",p.fname)
		fmt.Printf("\t")
		fmt.Printf("LAST NAME: %s",p.lname)
		fmt.Println()
	}
}