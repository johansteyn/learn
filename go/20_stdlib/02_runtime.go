package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go Standard Library: runtime")
	fmt.Println()

	fmt.Printf("runtime.GOOS: %s\n", runtime.GOOS)
	fmt.Printf("runtime.GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("runtime.Version(): %s\n", runtime.Version())
}
