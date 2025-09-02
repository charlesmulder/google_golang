package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sl := make([]int, 0, 3)
	var num string
	for {
		fmt.Print("Enter Number: ")
		fmt.Scanf("%s\n", &num)
		if num == "X" {
			break
		}
		iNum, _ := strconv.Atoi(num)
		sl = append(sl, iNum)
		sort.IntSlice.Sort(sl)
		fmt.Println(sl)
	}
}
