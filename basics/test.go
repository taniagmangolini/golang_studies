package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id  int
	ate int
}

func (p *Philo) eat() {
	for i := 0; i < 3; i++ {
		leftCS, rightCS := host(p.id)

		leftCS.Lock()
		rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.id+1)
		p.ate = p.ate + 1
		fmt.Printf("finishing eating %d\n", p.id+1)

		rightCS.Unlock()
		leftCS.Unlock()
	}

	wg.Done()
}

var CSticks = make([]*ChopS, 5)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func host(philo_id int) (*ChopS, *ChopS) {
	var left *ChopS
	var right *ChopS

	id1 := min(philo_id, (philo_id+1)%5)
	id2 := max(philo_id, (philo_id+1)%5)

	left = CSticks[id1]
	right = CSticks[id2]

	return left, right
}

var wg sync.WaitGroup

func main() {

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, 0}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

	wg.Wait()

	for i := 0; i < 5; i++ {
		fmt.Printf("Philosopher %d ate %d times\n", philos[i].id, philos[i].ate)
	}
}