package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
func main() {

	// Create slice with length=0 and capacity=3
	slice := make([]int, 0, 3)

	for {

		var input string

		// Accept user input
		fmt.Println("Please enter an integer to add to the slice:")
		_, err := fmt.Scanln(&input)

		if err != nil {
			log.Fatal(fmt.Errorf("Scanning input: %w", err))
		}

		// Check for exit command
		if strings.ToLower(input) == "x" {
			fmt.Println("Exiting")
			os.Exit(0)
		}

		// Convert string to int
		i, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(fmt.Errorf("Converting string input to int: %w", err))
		}

		// Append int to slice
		slice = append(slice, i)

		// Sort slice
		sort.Ints(slice)

		// Print slice
		fmt.Println(slice)
	}

}
