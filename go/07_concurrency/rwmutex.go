package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var mutex sync.RWMutex
var wg sync.WaitGroup
var readlock bool

func main() {
	fmt.Printf("Concurrency - rwmutex\n")
	fmt.Println()

	args := os.Args[1:]
	num := 5
	if len(args) >= 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid number: %s\n", args[0])
			os.Exit(1)
		}
		num = i
		if len(args) == 2 && args[1] == "r" {
			readlock = true
		}
	}

	multiplier := 1000
	wg.Add(num + num*multiplier)
	start := time.Now()
	fmt.Printf("Starting %d producers and %d consumers...\n", num, num*multiplier)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			produce()
		}()
		for j := 0; j < multiplier; j++ {
			go func() {
				defer wg.Done()
				consume()
			}()
		}
	}
	fmt.Println("Waiting...")
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Done.")
	fmt.Printf("Time taken: %v seconds\n", duration)
}

// Producing takes at full second
func produce() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("*** Producing...")
	pause(1000)
}

// Consuming takes only a millisecond
// and can use either the normal lock
// or the read lock, which will only
// block while producing
func consume() {
	if readlock {
		mutex.RLock()
		defer mutex.RUnlock()
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}
	fmt.Println("Consuming...")
	pause(1)
}

func pause(millis time.Duration) {
	time.Sleep(millis * time.Millisecond)
}
