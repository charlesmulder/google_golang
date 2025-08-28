package main

import (
	"fmt"
)

// s = 0.5 * a * t^2 + v0 * t + s0
func main() {

	var a, v, d float64
	fmt.Println("Enter acceleration:")
	fmt.Scanln(&a)
	fmt.Println("Enter velocity:")
	fmt.Scanln(&v)
	fmt.Println("Enter displacement:")
	fmt.Scanln(&d)

	fmt.Printf("a: %.2f\nv: %.2f\nd: %.2f\n", a, v, d)
}
