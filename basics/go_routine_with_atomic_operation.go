package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int32

	// Goroutine 1: Increment counter
	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddInt32(&counter, 1)
		}
	}()

	// Goroutine 2: Decrement counter
	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddInt32(&counter, -1)
		}
	}()

	// Give goroutines time to execute
	time.Sleep(1 * time.Second)

	// Print the final value of counter
	fmt.Println("Final counter value:", counter)
}
