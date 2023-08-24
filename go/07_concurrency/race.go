package main

import (
	"fmt"
	"time"
)

// % go build race.go
// % while true; do race; sleep 1; done
func main() {
	fmt.Printf("Concurrency - race\n")
	fmt.Println()

	n := 0
	go inc(&n)
	// By pausing and printing we allow enough time for the function call to complete,
	// thereby only looping once while printing 1
	counter := 0
	for ; n == 0; counter++ {
		// Without any pause, the loop is sometimes run more than once, while printing zero
		pause(1)
		// Without any pause or print, the loop is run thousands of times!
		fmt.Printf("#%d: %d\n", counter, n)
	}
	fmt.Printf("Looped %d times\n", counter)
}

func inc(n *int) {
	*n++
}

func pause(millis time.Duration) {
	time.Sleep(millis * time.Millisecond)
}
