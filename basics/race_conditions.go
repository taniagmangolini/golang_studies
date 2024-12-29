package main

import (
	"fmt"
	"time"
)

func main() {
	/* This program has a problem of race condition
	because we have two go routines performing operations
	over the same variable (counter) without any kind of lock.
	So, it can cause unexpected counter results.
	*/
	var counter int

	go func() {
		for i := 0; i < 100000; i++ {
			counter++
			fmt.Println("|G1", counter)
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			counter--
			fmt.Println("|G2", counter)
		}
	}()

	time.Sleep(3 * time.Second) //it gives time to  the go routines execute before main completes

	fmt.Println("Final counter value:", counter)
}
