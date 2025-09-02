package main

import "fmt"

func main() {
	fmt.Print("Hi! Please enter a float: ")

	var fNumber float32
	fmt.Scanln(&fNumber)

	var iNumber int = int(fNumber)

	fmt.Println("Truncated part is:", iNumber)

}
