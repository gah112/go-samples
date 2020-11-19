package main

import (
    "fmt"
    "strconv"
)

const PRECISION = 64

func GenDisplaceFn(ai float64, vi float64, si float64) func(float64) float64 {
	// Create the parameterized displacement function:
	fn := func(t float64) float64 {
		return 0.5 * ai * t * t + vi * t + si
	}
	
	return fn
}

func main() {
	var input string

	// Prompt the user for an initial position:
	fmt.Println("Please enter an initial position...")
	fmt.Scan(&input)
	s, err := strconv.ParseFloat(input, PRECISION)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Prompt the user for an initial velocity:
	fmt.Println("Please enter an initial velocity...")
	fmt.Scan(&input)
	v, err := strconv.ParseFloat(input, PRECISION)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Prompt the user for an initial acceleration:
	fmt.Println("Please enter an initial acceleration...")
	fmt.Scan(&input)
	a, err := strconv.ParseFloat(input, PRECISION)
	if err != nil {
		fmt.Println(err)
		return
	}

	fn := GenDisplaceFn(s, v, a)

	// Prompt the user for a time:
	fmt.Println("Please enter a time to evaluate the displacement function...")
	fmt.Scan(&input)
	t, err := strconv.ParseFloat(input, PRECISION)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Displacement [s(t)] = %f at t = %f", fn(t), t)
	}
}