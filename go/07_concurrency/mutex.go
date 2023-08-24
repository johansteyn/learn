package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var counter int
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	fmt.Printf("Concurrency - mutex\n")
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

	wg.Add(2 * num)
	fmt.Printf("Starting %d inc goroutines...\n", num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			inc()
		}()
	}
	fmt.Printf("Starting %d dec goroutines...\n", num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			dec()
		}()
	}
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Done.")
	// Counter should always end in zero
	fmt.Printf("Counter = %d\n", counter)
}

func inc() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("Incrementing %d...\n", counter)
	counter++
}

func dec() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("Decrementing %d...\n", counter)
	counter--
}
