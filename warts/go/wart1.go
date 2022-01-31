package main

import (
	"fmt"
)

func main() {
	var slice1 []int
	var slice2 = []int{}
	//fmt.Printf("slice1 == slice2? %t\n", slice1 == slice2)
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)
	fmt.Printf("len(slice1) = %d\n", len(slice1))
	fmt.Printf("len(slice2) = %d\n", len(slice2))
	fmt.Printf("slice1 == nil? %t\n", slice1 == nil)
	fmt.Printf("slice2 == nil? %t\n", slice2 == nil)
}

