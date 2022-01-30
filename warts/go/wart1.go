package main

import (
	"fmt"
)

func main() {
	var slice1 []int
	var slice2 = []int{}
	//fmt.Printf("slice1 == slice2? %t\n", slice1 == slice2)
	fmt.Printf("slice1 == nil? %t\n", slice1 == nil)
	fmt.Printf("slice2 == nil? %t\n", slice2 == nil)
	fmt.Printf("slice1=%v (%T)\n", slice1, slice1)
	fmt.Printf("slice2=%v (%T)\n", slice2, slice2)
}

