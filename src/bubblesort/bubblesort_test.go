package main

import (
	"testing"
	"slices"
	"math/rand"
)

func TestSwap(t *testing.T) {

	index := rand.Intn(9)
	numbers := make([]int, 10)
	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	nr1, nr2 := numbers[index], numbers[index+1]

	Swap(numbers, index)

	if numbers[index] != nr2 || numbers[index+1] != nr1 {
		t.Errorf("Numbers not swapped")
	}

}

func TestBubbleSort(t *testing.T) {
	x := []int{3,5,2,7,1}
	BubbleSort(x)

	if ! slices.Equal( x, []int{1,2,3,5,7} ) {
		t.Errorf("Numbers not sorted")
	}
}
