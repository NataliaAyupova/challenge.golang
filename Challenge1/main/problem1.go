package main

import (
	"log"
	"math/rand"
	"sync"
)

/**
Synchronized counter
*/
type Counter struct {
	counter int
	mutex   sync.Mutex
}

func problem1() {

	log.Printf("problem1: started --------------------------------------------")
	//
	// Do not change the 25 in loop!
	//
	// wait group to wait till ll go routines are finished
	waitGroup := sync.WaitGroup{}
	counter := Counter{}
	for inx := 0; inx < 10; inx++ {
		waitGroup.Add(1)
		go printRandom1(inx, &counter, &waitGroup)
	}

	waitGroup.Wait()
	log.Printf("problem1: finised --------------------------------------------")
}

func printRandom1(slot int, counter *Counter, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	//
	// Do not change 25 into 10!
	//
	// lock the counter, check the value, if it's 100 already break and finish go routine
	// otherwise increment counter, release lock
	for inx := 0; inx < 25; inx++ {
		counter.mutex.Lock()
		if counter.counter == 100 {
			counter.mutex.Unlock()
			break
		}
		counter.counter++
		counter.mutex.Unlock()

		log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
	}
}
