package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	nr      int
	c       chan bool
	leftCS  *ChopStick
	rightCS *ChopStick
}

const philoCount = 5
const mealCount = 3

var wg sync.WaitGroup
var onceDo sync.Once

func main() {

	// init 5 channels, one for each philo
	var c [5]chan bool
	for i := 0; i < 5; i++ {
		c[i] = make(chan bool)
	}

	// readiness channel
	cr := make(chan int)

	sticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopStick)
	}

	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{i + 1, c[i], sticks[i], sticks[(i+1)%5]}
	}

	for i := 0; i < philoCount; i++ {
		wg.Add(1)
		go philos[i].eat(cr)
	}

	wg.Add(1)
	go host(&c, cr)

	wg.Wait()
}

// deadlock mogelijkheid
func (p Philosopher) eat(cr chan int) {

	for i := 1; i < 4; i++ {

		<-p.c // wait for availability

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("philosopher", p.nr, " start eating", i)
		time.Sleep(time.Second)
		fmt.Println("philosopher", p.nr, " end eating", i)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		cr <- p.nr // signal readiness
	}
	wg.Done()
}

func host(c *[5]chan bool, cr chan int) {

	// start eating for first two philo's
	c[0] <- true
	c[2] <- true

	next := 3
	for i := 2; i < 15; i++ {
		<-cr // wait for philo's signal of readiness
		next = whosNext(next)
		c[next-1] <- true
	}

	<-cr
	<-cr

	fmt.Println("host done")
	wg.Done()
}

func whosNext(current int) int {
	next := (current + 2) % 5
	if next == 0 {
		next = 5
	}
	return next
}
