package main

import (
	"fmt"
	"slices"
	"sync"
)

// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 
// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
// echo -e "5\n3\n8\n1\n9\n2\n7\n4\n6\nq" | go run src/threadsort/threadsort.go
func main() {

	var numbers []int
	var number int
	var wg sync.WaitGroup

	fmt.Println("Enter integers one at a time (press Enter after each). Type any non-number to start sorting:")
	for {
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			break
		}
		numbers = append(numbers, number)
	}
	fmt.Println("You entered:")
	fmt.Println(numbers)
	fmt.Println("---")

	partitions := Partition(numbers, 4) // partition into 4 parts

	wg.Add(len(partitions))

	fmt.Println("Partitions")
	for _, p := range partitions {
		fmt.Println(p)
		go Sort(p, &wg)
	}
	fmt.Println("---")

	wg.Wait() // wait for sorting jobs to finish

	// access sorted partitions
	merged1 := Merge(partitions[0], partitions[1])
	merged2 := Merge(partitions[2], partitions[3])
	fmt.Println( "Sorted array: ")
	fmt.Println( Merge(merged1, merged2 ))
	fmt.Println("---")

}

func Merge( left []int, right []int ) []int {

	result := make([]int, 0, len(left)+len(right) )
	leftPtr, rightPtr := 0, 0

	for leftPtr < len(left) && rightPtr < len(right) {
		if left[leftPtr] <= right[rightPtr] {
			result = append(result, left[leftPtr])
			leftPtr++
		} else {
			result = append(result, right[rightPtr])
			rightPtr++
		}
	}

	// result = append(result, left[leftPtr:]...)
	for leftPtr < len(left) {
		result = append(result, left[leftPtr])
		leftPtr++
	}

	//result = append(result, right[rightPtr:]...)
	for rightPtr < len(right) {
		result = append(result, right[rightPtr])
		rightPtr++
	}

	return result
}

func Partition( src []int,  parts int ) [][]int {

	length := len(src)
	quotient := length / parts // Items per partition
	remainder := length % parts // Remainder is always < divisor

	target := make([][]int, parts)  
	start := 0

	for i := range parts {
		size := quotient
		if i < remainder {
			size++  // Distribute remainder to first few partitions
		}

		end := start + size
		if start < length {
			target[i] = make([]int, size)
			copy(target[i], src[start:end])
		}
		start = end
	}

	return target
}

func Sort( numbers []int, wg *sync.WaitGroup) {
	if( wg != nil ) {
		defer wg.Done()
	}
	slices.Sort(numbers)
}

