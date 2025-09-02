package main

import (
	"fmt"
	"math"
)

func main() {

	// acceleration, initial velocity, initial displacement, time
	var a, v0, s0, t float64 // time

	fmt.Println("Enter acceleration:")
	fmt.Scanln(&a)
	fmt.Println("Enter initial velocity:")
	fmt.Scanln(&v0)
	fmt.Println("Enter initial displacement:")
	fmt.Scanln(&s0)
	
	for {
		fmt.Println("Enter time (Ctrl+C to quit):")
		fmt.Scanln(&t)
		Disp1 := GenDisplaceFn( a, v0, s0 )
		fmt.Printf("Displacement is: %.2f\n\n", Disp1(t))
	}

}

func GenDisplaceFn( a, v0, s0 float64 ) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5 * a * math.Pow(t, 2) + v0 * t + s0
	}
}
