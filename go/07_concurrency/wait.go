package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("Concurrency - wait\n")
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
	}

	// Add the number of goroutines we'll be running
	// NOTE: We could instead call wg.Add(1) before each individual goroutine call
	wg.Add(num)
	fmt.Printf("Starting %d goroutines...\n", num)
	start := time.Now()
	for i := num; i > 0; i-- {
		go foo(time.Duration(i % 10))
	}
	fmt.Println("Waiting...")
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Done.")
	fmt.Printf("Time taken: %v seconds\n", duration)
}

func foo(seconds time.Duration) {
	defer wg.Done()
	fmt.Printf("Running foo for %d second(s)...\n", seconds)
	pause(seconds * 1000)
}

func pause(millis time.Duration) {
	time.Sleep(millis * time.Millisecond)
}
