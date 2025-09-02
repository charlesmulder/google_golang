package main

import "fmt"

func Swap(sliceToSort []int, j int) {
	temp := sliceToSort[j]
	sliceToSort[j] = sliceToSort[j+1]
	sliceToSort[j+1] = temp
}

func BubbleSort(sliceToSort []int) {
	//buble sort two loops
	for i := 0; i < len(sliceToSort)-1; i++ {
		for j := 0; j < len(sliceToSort)-i-1; j++ {
			if sliceToSort[j] > sliceToSort[j+1] {
				Swap(sliceToSort, j)
			}

		}
	}

}

func main() {

	sliceToSort := []int{}
	var readNumber int
	// Prompt type in a sequence of up to 10 integers.

	for i := 0; i < 10; i++ {
		fmt.Printf("Enter value [%d out of 10] :", i+1)
		fmt.Scan(&readNumber)
		sliceToSort = append(sliceToSort, readNumber)
	}

	BubbleSort(sliceToSort)
	fmt.Printf("\n\nSorted slice :")
	for i := range sliceToSort {
		fmt.Printf("%d ", sliceToSort[i])
	}
	fmt.Printf("\n")
}
