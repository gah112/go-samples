package main

import (
    "fmt"
    "sort"
    "strconv"
)

const STOP_CHARACTER = "x"

func main() {
    var input string
    var array []int

    for true {
        // Prompt the user for an integer:
        fmt.Println("Enter an integer...")

        // Get the user's input:
        fmt.Scan(&input)
        
        // Break if the user enters the stop character:
        if input == STOP_CHARACTER {
            fmt.Println("Stopping process...")
            return
        }

        // Try conversion of the input string to an integer:
        num, err := strconv.Atoi(input)

        // If conversion results in an error, throw a message:
        if (err != nil) {
            fmt.Println("ERROR: Invalid input - you must provide an integer!")
        } 
        else {
            // Add the number to the slice:
            array = append(array,num)
            
            // Sort the array:
            sort.Ints(array)

            // Print the output:
            fmt.Println(array)
        }
    }
}