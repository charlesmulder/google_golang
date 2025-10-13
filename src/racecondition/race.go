package main

import (
	"fmt"
	"time"
)

// Most of the time the program prints 0, but on occasion it prints 2. Since the instructions
// are interleaved at the machine level it is also possible to print 1.
//
// The program prints 0 when Goroutine 2 executes before Goroutine 1 starts
// The program prints 1 when Goroutine 2 executes after x = 1 but before x = x + 1
// The program prints 2 when Goroutine 2 executes after Goroutine 1 completes both operations
//
// The race condition occurs because both Goroutines access the shared variable x simultaneously without synchronization.
// The timing of when each Goroutine executes determines the output.
// The output of the program is non-deterministic.
func main() {
	var x int = 0

	// Assign and increment function
	go func() {
		x = 1
		x = x + 1
	}()

	// Print function
	go func() {
		fmt.Println(x)
	}()

	time.Sleep(200 * time.Millisecond)
}
