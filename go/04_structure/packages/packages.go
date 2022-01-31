package main

import (
	"fmt"
	"packages/mypackage" // Note: use forward slashes, not dots
)

func main() {
	fmt.Println("Packages")
	fmt.Println()
	mypackage.Foo()
	fmt.Println("mypackage.MeaningOfLife:", mypackage.MeaningOfLife)
}

