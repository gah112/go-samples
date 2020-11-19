package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	var unsorted []int
	
	// Prompt the user to enter a comma-separated list of integers:
	fmt.Println("Enter a comma-separated list of up to 10 integers...")

	// Scan for the input string:
	fmt.Scan(&input)

	// Split the input string:
	inputs := strings.Split(input, ",")

	for j, str := range inputs {
		// Convert the string to an integer:
		num, err := strconv.Atoi(str)

		// Append the integers to the list:
		if err != nil {
			fmt.Printf("WARNING: %s is not a valid integer and will be discarded. Continuing.\n",str)
		} else if len(unsorted) > 10 {
			fmt.Printf("WARNING: More than 10 integers have been provided and the values %s will be discarded. Continuing.\n",inputs[j:])
			break
		} else {
			unsorted = append(unsorted,num)
		}
	}

	// Sort the array using Bubble Sort:
	sorted := BubbleSort(unsorted)

	// Print the integers:
	fmt.Println(sorted)
}

func Swap(list []int, index int) {
	temp := list[index]
	list[index] = list[index + 1]
	list[index + 1] = temp
}

func BubbleSort(list []int) []int {
	// Perform the sort by swapping integers:
	swapped := true
	for swapped {
		swapped = false
		for j := 0; j < len(list) - 1; j++ {
			if list[j] > list[j + 1] {
				Swap(list,j)
				swapped = true
			}
		}
	}

	// Return the sorted list:
	return list
}