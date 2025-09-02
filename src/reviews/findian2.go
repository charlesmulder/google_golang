package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const figure rune = 'a'

	// To Take full string as input with variable whitespaces to handle tests like {I d skd a efju N}, so am using bufio.NewScanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	a := strings.ToLower(str)

	if strings.HasPrefix(a, "i") && strings.HasSuffix(a, "n") {
		for _, val := range a {
			if val == figure {
				fmt.Println("Found!")
				return
			}
		}
	}
	fmt.Println("Not Found!")
}
