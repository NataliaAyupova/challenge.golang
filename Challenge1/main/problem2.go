package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

/**
Synchronized ticker
*/
type SynchTicker struct {
	ticker *time.Ticker
	mutex  sync.Mutex
}

func problem2() {

	log.Printf("problem2: started --------------------------------------------")
	// wait group to wait till ll go routines are finished
	waitGroup := sync.WaitGroup{}
	syncTicker := SynchTicker{ticker: time.NewTicker(time.Second)}

	for inx := 0; inx < 10; inx++ {
		waitGroup.Add(1)
		go printRandom2(inx, &waitGroup, &syncTicker)
	}

	waitGroup.Wait()

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int, waitGroup *sync.WaitGroup, syncTicker *SynchTicker) {
	defer waitGroup.Done()

	for inx := 0; inx < 10; inx++ {
		waitOnTicker(syncTicker)
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
	}
}

/**
locks the ticker and waits till it ticks, after that releases the lock so that other go routines can proceed
*/
func waitOnTicker(syncTicker *SynchTicker) {
	defer syncTicker.mutex.Unlock()
	syncTicker.mutex.Lock()
	for {
		select {
		case <-syncTicker.ticker.C:
			return
		}
	}
}
