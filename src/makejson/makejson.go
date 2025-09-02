package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
func main() {

	profile := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter your name:")
	scanner.Scan()
	profile["name"] = scanner.Text()
	fmt.Println("Please enter your address:")
	scanner.Scan()
	profile["address"] = scanner.Text()
	json, err := json.Marshal(profile)
	if err != nil {
		log.Fatal(fmt.Errorf("Converting string input to int: %w", err))
	}
	fmt.Println(string(json))
}
