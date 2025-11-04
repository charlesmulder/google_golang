/**
 * Implement the dining philosopher’s problem with the following constraints/modifications. 
 * There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
 * Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
 * The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
 * In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
 * The host allows no more than 2 philosophers to eat concurrently.
 * Each philosopher is numbered, 1 through 5.
 * When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
 * When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
 */

package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
	"sync/atomic"
)

type Chopstick struct {}

type Philosopher struct {
	id int
	chopsticks [2]chan Chopstick
	timesEaten uint32
	isEating bool
	mu sync.Mutex
}

type EatingPermit struct {}

// Ensure all philosophers eat 3 times
func hasFinishedEating( philosophers []*Philosopher ) bool {
	for _, p := range philosophers {
		if atomic.LoadUint32(&p.timesEaten) < 3 {
			return false
		}
	}
	return true
}

func (p *Philosopher) eat( permits chan EatingPermit ) {

	randNum := rand.Intn(2)
	// Try to get chopsticks without blocking
	select {
	case <-p.chopsticks[randNum]:
		select {
		case <-p.chopsticks[1-randNum]:
			// Got both chopsticks!
			<-permits // Now get permit
		default:
			// Couldn't get second chopstick
			p.chopsticks[randNum] <- Chopstick{} // Return first
			return // Give up
		}
	default:
		return // Couldn't get first chopstick, give up
	}

	p.mu.Lock()
	p.isEating = true
	p.mu.Unlock()
	fmt.Printf("starting to eat %d\n", p.id)
	time.Sleep(time.Millisecond * 500)
	atomic.AddUint32(&p.timesEaten, 1)
	// return chopsticks to channels
	p.chopsticks[randNum] <- Chopstick{}
	p.chopsticks[1 - randNum ] <- Chopstick{}
	// return eating permit
	permits<-EatingPermit{}
	p.mu.Lock()
	p.isEating = false
	p.mu.Unlock()
	fmt.Printf("finishing eating %d (times eaten: %d)\n", p.id, atomic.LoadUint32(&p.timesEaten))
}

func host( wg *sync.WaitGroup, philosophers []*Philosopher ) {

	defer wg.Done()

	// Only allow 2 philosophers to eat at a time
	permits := make(chan EatingPermit, 2)
	// Add 2 eating permits
	permits <- EatingPermit{}
	permits <- EatingPermit{}

	for ! hasFinishedEating( philosophers ) {

		time.Sleep(time.Millisecond * 10) // prevent busy loop

		for _, p := range philosophers {
			p.mu.Lock()
			canEat := ! p.isEating 
			p.mu.Unlock()
			if canEat && atomic.LoadUint32(&p.timesEaten) < 3 {
				go p.eat(permits)
			}
		}
	}

}

func main() {
	var wg sync.WaitGroup

	// init 5 chopsticks
	tableChopsticks := make([]chan Chopstick, 5, 5)
	for i := range tableChopsticks {
		tableChopsticks[i] = make(chan Chopstick, 1)
		tableChopsticks[i] <- Chopstick{} // add chopstick to channel
	}

	// init 5 philosophers
	var philosophers = make([]*Philosopher, 5, 5)
	for i := range 5 {
		philosophers[i] = &Philosopher{
			id: i+1,
			chopsticks: [2]chan Chopstick{ tableChopsticks[i], tableChopsticks[(i+1)%5] },
			timesEaten: 0,
			isEating: false,
		}
	}
	wg.Add(1)
	go host( &wg, philosophers )
	wg.Wait()
}
