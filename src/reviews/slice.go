package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
	fmt.Println("Entare an integer to store or X to quit \n")
	// Create a slice of length0, capacity 3 
	nums := make([]int, 0, 3)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter an integer (or X to exit): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if strings.EqualFold(input, "X") {
            fmt.Println("You entered X. Exiting program.")
            break
        }
        num, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Invalid input. Please enter an integer or X.")
            continue
        }
        nums = append(nums, num) // Add to slice
        sort.Ints(nums)          // Sort slice

        fmt.Println("Sorted slice:", nums)

	}
}
