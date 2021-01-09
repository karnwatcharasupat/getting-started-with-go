package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	number          int
	leftCS, rightCS *ChopStick
}

var EATING = 2
var FINISHED = 0
var DONE = 4

var REQUEST = 1
var DENIED = 5
var APPROVED = 3

func main() {
	chopsticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(ChopStick)
	}

	philos := make([]*Philosopher, 5) // create slice

	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{
			i,
			chopsticks[i],
			chopsticks[(i+1)%5],
		}
	}

	var chans [5]chan int
	for i := range chans {
		chans[i] = make(chan int)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go hostRoutine(chans, &wg, chopsticks)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg, chans[i])
	}

	wg.Wait()
}

func sumBool(arr []bool) int {
	n := 0

	for _, a := range arr {
		if a {
			n++
		}
	}

	return n
}

func hostRoutine(channels [5]chan int, wg *sync.WaitGroup, cs []*ChopStick) {
	// fmt.Println("host started")
	allDone := false
	dones := make([]bool, 5)
	eating := make([]bool, 5)

	for !allDone {
		// fmt.Println("host checking")
		for i, c := range channels {
			select {
			case msg := <-c:
				switch msg {
				case DONE:
					dones[i] = true
					eating[i] = false
				case REQUEST:
					// fmt.Println("host received request from ", i)
					if sumBool(eating) < 2 && !eating[(i+6)%5] && !eating[(i+1)%5] {
						// fmt.Println("host approved request from ", i)
						c <- APPROVED
						eating[i] = true
					} else {
						c <- DENIED
					}
				case FINISHED:
					eating[i] = false
				}
			default:
				continue
			}
		}

		allDone = true
		for _, b := range dones {
			allDone = allDone && b
		}
	}

	wg.Done()

}

func (p Philosopher) eat(wg *sync.WaitGroup, c chan int) {
	// fmt.Println("philo created", p.number)
	ready := false

	for i := 0; i < 3; i++ {
		ready = false
		c <- REQUEST
		// fmt.Println("requesting to eat ", p.number)

		for !ready {
			select {
			case msg := <-c:
				if msg == APPROVED {
					ready = true
				} else {
					c <- REQUEST
				}
			default:
				continue
			}
		}

		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("Philosopher", p.number, "is starting to eat.")

		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Microsecond)

		fmt.Println("Philosopher", p.number, "is finishing eating.")
		c <- FINISHED

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		r = rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Microsecond)
	}

	fmt.Println("Philosopher", p.number, "is done eating 3 times")

	c <- DONE
	wg.Done()
}
