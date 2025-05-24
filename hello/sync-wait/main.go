package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)
	// some task is happening
	fmt.Printf("worker %d stopped\n", i)
}

// worker(1)
// worker(2)
// worker(3)

func main() {
	// fmt.Println("learning about syncWaitGroup function.")

	var wg sync.WaitGroup

	// start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment the wait sync group.
		go worker(i, &wg)
	}

	// wait for all the workers to finish
	wg.Wait()
	fmt.Println("worker task is completed.")

}
