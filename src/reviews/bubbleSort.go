package main

import (
	"fmt"
)

func Swap(slice []int, i int) {
	slice[i],slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int) {
	n :=len(slice)
	for i:=0; i<n-1; i++ {
		for j :=0;j <n-i-1;j++ {
			if slice[j] >slice[j+1] {
				Swap(slice,j)
			}
		}
	}

}

func main() {
	var input int
	var nums []int

	fmt.Println("Enter the interger[max 10 numbers and non-interger to stop]: ")
	for len(nums) <10 {
		_,err :=fmt.Scan(&input)
		if err!=nil {
			break
		}
		nums = append(nums, input)
	}
	
	BubbleSort(nums)

	fmt.Println("Sorted Nums:",nums)

}