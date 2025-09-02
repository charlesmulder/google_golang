package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/
func main() {

	var fileName string

	type Person struct { // Person struct
		fname [20]byte // Each field will be a string of size 20 (characters)
		lname [20]byte
	}

	var people []Person // People slice

	// Prompt the user for the name of the text file
	fmt.Println("Please enter file name:")
	_, err := fmt.Scanln(&fileName)
	if err != nil {
		log.Fatal(fmt.Errorf("Error scanning file name: %w", err))
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file: %w", err))
	}
	defer file.Close()

	// Your program will successively read each line of the text file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var person Person // create a struct which contains the first and last names found in the file
		copy(person.fname[:], parts[0])
		copy(person.lname[:], parts[1])
		people = append(people, person) // Each struct created will be added to a slice
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Program should iterate through your slice of structs and print the first and last names found in each struct.
	for _, p := range people {
		fmt.Printf("%s %s\n", string(p.fname[:]), string(p.lname[:]))
	}

}
