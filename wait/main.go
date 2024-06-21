package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(n int) error {
	println("working...")
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	if n == 2 {
		return fmt.Errorf("error oops")
	}

	println("done")
	return nil
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		// Increase the WaitGroup counter
		wg.Add(1)
		go func(n int) {
			// Decrease the WaitGroup counter when the goroutine completes
			defer wg.Done()
			// Call the worker function
			err := worker(n)
			// Check for errors
			if err != nil {
				println(err.Error())
			}
		}(i)
	}

	// Wait for all workers to finish
	wg.Wait()
}
