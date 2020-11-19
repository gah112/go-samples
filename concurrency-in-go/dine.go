package main

import (
	"fmt"
    "sync"
)

const NUM_PHILOSOPHERS = 5

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup

func main() {
	// Create the Channel:
	permission := make(chan int)

	// Create the chopsticks:
	chopsticks := make([]*Chopstick, NUM_PHILOSOPHERS)
	for j := 0; j < NUM_PHILOSOPHERS; j++ {
		chopsticks[j] = new(Chopstick)
	}

	// Create the philosophers:
	philosophers := make([]*Philosopher, NUM_PHILOSOPHERS)
	for j:= 0; j < NUM_PHILOSOPHERS; j++ {
		philosophers[j] = &Philosopher{j, 0, chopsticks[j], chopsticks[((j+1) % 5)]}
	}

	// Start the Host task:
	fmt.Println("Host has entered...")
	go host(philosophers, permission, &wg2)

	// Start the philosopher tasks:
	for j := range philosophers {
		// Add a task to the Wait Group:
		wg1.Add(1)

		// Start the Philosopher task:
		fmt.Println("Philosopher", philosophers[j].seat, "has entered...")
		go philosophers[j].eat(&wg1, &wg2, permission)
	}
	wg1.Wait()
}

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	seat int
	count int
	left *Chopstick
	right *Chopstick
}

func (philosopher Philosopher) eat(wg1 *sync.WaitGroup, wg2 *sync.WaitGroup, permission chan int) {
	for philosopher.count < 3 {
		for {
			// Ask for permission from the host:
			p := philosopher.ask(permission)

			if(p == philosopher.seat) {
				// Print that permission has been granted:
				fmt.Println("Philosopher", philosopher.seat, " has received permission to eat.")

				// Increment the number of times the philosopher has eaten:
				philosopher.count = philosopher.count + 1

				// Lock the chopsticks:
				philosopher.left.Lock()
				philosopher.right.Lock()
				fmt.Println("Philosopher", philosopher.seat, "has started eating.")

				// Release the chopsticks:
				philosopher.left.Unlock()
				philosopher.right.Unlock()
				fmt.Println("Philosopher", philosopher.seat, "has finished eating.")
				wg2.Done()
				break
			} else {
				// Print that permission has been denied:
				fmt.Println("Host has denied permission to Philosopher", philosopher.seat)
				permission <- p
			}
		}
	}
	wg1.Done()
}

func (philosopher Philosopher) ask(c chan int) int {
	a := <- c
	return a
}

func host(philosophers []*Philosopher, permission chan int, wg *sync.WaitGroup) {
	for j := 0; j < 8; j++ {
		// Add some tasks to the Wait Group:
		wg.Add(2)

		// Grant permission to the 1st philosopher:
		permission <- j % NUM_PHILOSOPHERS

		// Grant permission to the 2nd philosopher:
		if(j < 7) {
			permission <- (j+3) % NUM_PHILOSOPHERS
		}
		wg.Wait()
	}
}