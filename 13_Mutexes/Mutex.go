package main

import (
	"fmt"
	"sync"
	"time"
)

// Go's standard library provides a built-in implementation of a mutex with the sync.Mutex type and its two methods:
// .Lock()
// .Unlock()

type safeCounter struct {
	counts map[string]int
	mux    *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.slowVal(key)
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

//-------------------------------------------------------------------------------------------------

//RW MUTEX
// The sync.RWMutex also has these methods for concurrent reads:
// RLock()
// RUnlock()

type safeCounter2 struct {
	counts map[string]int
	mux    *sync.RWMutex
}

func (sc *safeCounter2) inc2(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement2(key)
}

func (sc *safeCounter2) val2(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key]
}

func (sc *safeCounter2) slowIncrement2(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

//-------------------------------------------------------------------------------------------------

func main() {
	// sc := &safeCounter{
	// 	counts: make(map[string]int),
	// 	mux:    &sync.Mutex{},
	// }

	// // Create a WaitGroup to wait for goroutines to finish
	// var wg sync.WaitGroup
	// keys := []string{"a", "b", "c"}

	// // Increment counters concurrently
	// for _, key := range keys {
	// 	wg.Add(1)
	// 	go func(key string) {
	// 		defer wg.Done()
	// 		for i := 0; i < 1000; i++ {
	// 			sc.inc(key)
	// 		}
	// 	}(key)
	// }

	// wg.Wait()

	// // Retrieve and print final counter values
	// for _, key := range keys {
	// 	fmt.Printf("Key: %s, Value: %d\n", key, sc.val(key))
	// }

	//Rw MUTEX
	sc := &safeCounter2{
		counts: make(map[string]int),
		mux:    &sync.RWMutex{},
	}

	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup
	keys := []string{"x", "y", "z"}

	// Increment counters concurrently
	for _, key := range keys {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				sc.inc2(key)
			}
		}(key)
	}

	wg.Wait()

	// Retrieve and print final counter values
	for _, key := range keys {
		fmt.Printf("Key: %s, Value: %d\n", key, sc.val2(key))
	}
}
