package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Controls the amount of parallelism:
const NUM_WORKERS = 4

func main() {
	var wg sync.WaitGroup
	var input string
	var unsorted []int
	var sorted []int
	var merged []int
	var list [][]int
	
	// Prompt the user to enter a comma-separated list of integers:
	fmt.Println("Enter a comma-separated list of up to 10 integers...")

	// Scan for the input string:
	fmt.Scan(&input)

	// Split the input string:
	inputs := strings.Split(input, ",")

	for j := range inputs {
		// Convert the string to an integer:
		num, err := strconv.Atoi(inputs[j])

		// Append the integers to the list:
		if err != nil {
			fmt.Printf("WARNING: %s is not a valid integer and will be discarded. Continuing.\n",inputs[j])
		} else {
			unsorted = append(unsorted,num)
		}
	}

	// Get the length of the slice:
	l := len(unsorted)

	// Create the channel for the workers:
	c := make (chan []int)

	// Set up the wait group:
	wg.Add(NUM_WORKERS)

	// Add each worker to the wait group:
	for j := 0; j < NUM_WORKERS; j++ {
		k := j + 1
		go worker(unsorted[(j*l/4):(k*l/4)], &wg, c)
	}

	// Get the worker results from the channel:
	for j := 0; j < NUM_WORKERS; j++ {
		sorted = <- c
		list = append(list,sorted)
	}

	// Await completion by all workers:
	wg.Wait()

	// Merge the worker results:
	for j := range list {
		merged = merge(merged,list[j])
	}

	// Print the merged list:
	fmt.Println(merged)
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
	return list
}

func worker(list []int, wg *sync.WaitGroup, c chan []int) {

	// Sort the list:
	sorted := BubbleSort(list)

	// Send the list to the channel:
	c <- sorted

	// Print the sorted list:
	fmt.Println(sorted)
	wg.Done()
}

func merge(a []int, b []int) []int {
	var res []int

	// Compare elements to merge the results:
	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			// Add the element from a to result:
			res = append(res,a[0])

			// Remove the element from a:
			a = a[1:len(a)]
		} else {
			// Add the element from b to result:
			res = append(res,b[0])

			// Remove the element from b:
			b = b[1:len(b)]
		}
	}	

	// Add remaining elements from a:
	for k := range a {
		res = append(res,a[k])
	}

	// Add remaining elements from b:
	for k := range b {
		res = append(res,b[k])
	}

	return res
}