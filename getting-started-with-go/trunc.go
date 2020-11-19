package main

import "fmt"

func main() {
	// Instantiate the float:
	var number float64

	// Prompt the user to enter a float:
	fmt.Scan(&number)

	// Print the truncated float as an integer:
	fmt.Println(int64(number))
}