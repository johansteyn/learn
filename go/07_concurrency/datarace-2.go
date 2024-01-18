package main

import (
	"fmt"
	"sync"
)

// https://go.dev/doc/articles/race_detector
// Race on loop counter
func main() {
	var wg sync.WaitGroup

	// This loop unexpectedly prints 55555 (or something like 55455) because
	// the iteration completes before any of the threads run the goroutine
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()

	// This loop correctly prints 01234 because we make a copy of i before any threads run the goroutine
	// Note, however, that it can print the digits in any order, eg: 42301 because threads run independently
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Print(j)
			wg.Done()
		}(i) // The i passed here is copied to the j parameter
	}
	wg.Wait()
	fmt.Println()
}
