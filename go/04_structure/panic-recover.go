package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Panic/recover (TODO...)")

	// https://www.youtube.com/watch?v=0c-1KJwSMCw
  // @7:00
	maybeGoexit()
	fmt.Println("Did not exit!")

}

func maybeGoexit() {
	fmt.Println("maybeGoexit...")
	defer func() {
		fmt.Println("Deferred function...")
	}()
	defer func() {
		fmt.Println(recover())
	}()
	defer panic("Exit cancelled...")
	runtime.Goexit()
}

