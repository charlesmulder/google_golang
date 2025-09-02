package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

/**
 * The program should print “Found!” if the entered string starts with 
 * the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
 */
func main() {

	var result bool
	var text string
	var scanner *bufio.Scanner

	fmt.Print("Hi! Please enter a string: ")

	scanner = bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text = scanner.Text()

	result = Findian(text)

	if result == true {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func Findian( haystack string ) bool {
	return StartsWithI( haystack ) && EndsWithN( haystack ) && ContainsA( haystack )
}

func StartsWithI( haystack string ) bool {
	return strings.HasPrefix( haystack, "i") || strings.HasPrefix( haystack, "I" )
}

func EndsWithN( haystack string ) bool {
	return strings.HasSuffix( haystack, "n") || strings.HasSuffix( haystack, "N" )
}

func ContainsA( haystack string ) bool {
	return strings.Contains( haystack, "a") || strings.Contains( haystack, "A" )
}
