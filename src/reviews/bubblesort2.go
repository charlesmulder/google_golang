package main

import (
	"fmt"
	"log"
)

func main() {

	seq := make([]int, 0, 10)

	fmt.Println("Enter a sequence of up to 10 integers. Press enter to stop.")
	for i := 10; i > 0; i-- {
		fmt.Printf("Enter an integer (%d left):\n", i)
		var nr int
		count, err := fmt.Scanln(&nr)

		// End of sequence
		if count != 1 {
			break
		}
		if err != nil {
			log.Fatal(fmt.Errorf("Error scanning integers: %w", err))
		}
		seq = append( seq, nr )
	}
	BubbleSort(seq)
	fmt.Println(seq)
}

func Swap( x []int, i int) {
	x[i], x[i+1] = x[i+1], x[i]
}

func BubbleSort( x []int ) {
	for i := len(x)-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if x[j] > x[j+1] {
				Swap(x, j)
			}
		}
	}
}
