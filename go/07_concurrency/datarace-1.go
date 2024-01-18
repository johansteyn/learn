package main

import (
	"fmt"
)

// https://go.dev/doc/articles/race_detector
// Introduction example
// Haven't figured it out yet - I don't see the conflicting access...
// One thread accesses the memory allocated to m["1"]
// The other thread accesses the memory allocated to m["1"]
// So.where's the conflict?
func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
