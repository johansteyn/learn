package main

import (
	"fmt"
	"math/rand"
	"testlist/datastructures"
	"time"
)

func main() {
	fmt.Println("Linked List")
	fmt.Println()

	list := datastructures.New()
	size := datastructures.GetTemperature()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		list.Add1(rand.Intn(100))
	}
	fmt.Printf("list: %s\n", list.String())
	fmt.Printf("list size: %d\n", list.Size())
	// Deliberately going out of bounds here
	for i := 0; i < size+1; i++ {
		value, err := list.Get(i)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("list.Get(%d): %d\n", i, value)
	}
}
