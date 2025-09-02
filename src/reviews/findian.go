package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputString string
	fmt.Println("Enter a string")
	fmt.Scan(&inputString)
	// Since we want ian and Ian to match, convert to lower case
	str := strings.ToLower(inputString)
	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		fmt.Println("Found!")
	}else{
		fmt.Println("Not found")
	}
}
