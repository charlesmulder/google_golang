
package main

import (
	"testing"
	"fmt"
	"math/rand"
	"slices"
	"github.com/stretchr/testify/assert"
)

func randRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func TestPartition(t *testing.T) {

	randNr := randRange(4,21) 
	numbers := rand.Perm( randNr )
	length := len(numbers)
	quotient := length / 4
	remainder := length % 4

	partitions := Partition(numbers, 4)

	fmt.Println(partitions)

	assert.Equal(t, len(partitions), 4)

	start := 0
	end := 0 
	for i := range 4 {
		end += quotient
		if i < remainder {
			end++
		}
		assert.Equal(t, len(partitions[i]), end-start)
		for j := range end-start {
			assert.True(t, slices.Contains(partitions[i], numbers[start+j] ) )
		}
		start = end
	}

}

