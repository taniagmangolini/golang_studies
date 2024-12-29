package main

import (
	"fmt"
	"sync"
)

var value int = 0

func showMsg(wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// both routines run concurrently
	go showMsg(&wg, "Go 1")
	go showMsg(&wg, "Go 2")

	wg.Wait()
	fmt.Println("End of program.")
}