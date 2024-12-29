package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	// Increment function
	increment := func() {
		defer wg.Done()
		mu.Lock()   // Acquire the lock
		counter++   // Critical section
		mu.Unlock() // Release the lock
	}

	wg.Add(2) // Two goroutines
	go increment()
	go increment()

	wg.Wait()
	fmt.Println("Final Counter Value:", counter)
}
