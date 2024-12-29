package main

import (
	"fmt"
	"sync"
)

const (
	MaxEatingPhilosophers = 2
    MaxEatingsByPhilosopher = 3
)

var wg sync.WaitGroup

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	eatCounter      int
}

func (p *Philo) eat(philoNumber int) {
	defer wg.Done()

	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Println("starting to eat", philoNumber)
	p.eatCounter++
	FinishedToEat <- philoNumber
	fmt.Println("finish to eat", philoNumber)

	p.rightCS.Unlock()
	p.leftCS.Unlock()
}

var AskToEat chan int
var FinishedToEat chan int
var eating []int
var mu sync.Mutex // Mutex for control eating edition

func removeElement(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func host() {
	for {
		select {
        // waits for requests to eat by the philosophers
		case philoNumber := <-AskToEat:
			mu.Lock()
            // only two different philosophers can eat concurrently
			if len(eating) < MaxEatingPhilosophers && !contains(eating, philoNumber) {
                // if the philosopher is allowed to eat, put him in the eating slice
				eating = append(eating, philoNumber)
				mu.Unlock()
				wg.Add(1)
                // So, allow philosopher to eat!
				go philos[philoNumber].eat(philoNumber)
			} else {
				mu.Unlock()
			}
        // waits for alerts of finishing eat by the philosophers
		case philoNumber := <-FinishedToEat:
			mu.Lock()
			eating = removeElement(eating, philoNumber)
			mu.Unlock()
		}
	}
}

var philos []*Philo

func main() {
    // create channels to request eating and to alert finish eating
	AskToEat = make(chan int, 2)
	FinishedToEat = make(chan int, 2)

    // create chopsticks 
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

    //create philosophers
	philos = make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], 0}
	}

    // Host controls the requests for eating and allow philosofers to eat.
	go host()

    // Philosofers ask to eat.
    // Each philosopher can eat only three times.
	for {
		allDone := true
		for i := 0; i < 5; i++ {
			if philos[i].eatCounter < MaxEatingsByPhilosopher {
				allDone = false
				AskToEat <- i
			}
		}
		if allDone {
			break
		}
	}

	wg.Wait()
	fmt.Println("All philosophers have finished eating.")
}
