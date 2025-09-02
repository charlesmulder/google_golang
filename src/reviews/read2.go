package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text
file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and
lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively
read each line of the text file and create a struct which contains the first and last names found
in the file. Each struct created will be added to a slice, and after all lines have been read
from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs
 and print the first and last names found in each struct.

Submit your source code for the program, “read.go”.
*/

type Name struct {
	fname [20]rune
	lname [20]rune
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a filename: ")
	s.Scan()
	fname := s.Text()
	fmt.Println("Opening file ", fname)
	infile, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	var names []Name
	inbuf := bufio.NewScanner(infile)
	for inbuf.Scan() {
		var n Name
		var fname, lname string
		np, err := fmt.Sscanf(inbuf.Text(), "%s %s\n", &fname, &lname)
		if err != nil {
			fmt.Println("Problem scanning name:", err)
		}
		fmt.Println("Parsed", np, "fields, name", fname, lname)
		for i := 0; i < 20; i++ {
			if len(fname) > i {
				n.fname[i] = rune(fname[i])
			}
			if len(lname) > i {
				n.lname[i] = rune(lname[i])
			}
		}
		names = append(names, n)
	}
	fmt.Println("Reached EOF")
	fmt.Println()
	for _, name := range names {
		fmt.Println(string(name.fname[:]), string(name.lname[:]))
	}
}
