package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mutex sync.Mutex
	value int
}

var wg sync.WaitGroup

func main() {
	fmt.Printf("Concurrency - deadlock\n")
	fmt.Println()

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
	fmt.Println("Done")
}

func printSum(a, b *value) {
	defer wg.Done()
	a.mutex.Lock()
	defer a.mutex.Unlock()
	pause(200)
	b.mutex.Lock()
	defer b.mutex.Unlock()
	sum := a.value + b.value
	fmt.Printf("%d + %d = %d\n", a.value, b.value, sum)
}

func pause(millis time.Duration) {
	time.Sleep(millis * time.Millisecond)
}
