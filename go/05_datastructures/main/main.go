package main

import (
	"datastructures"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Linked List")
	fmt.Println()
	list := datastructures.NewList()
	size := 10
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		value := rand.Intn(100) + 1
		if value%2 == 0 {
			fmt.Printf("Adding %d..\n", value)
			list.Add(value)
		} else {
			index := 0
			if list.Size() > 0 {
				index = value % list.Size()
			}
			fmt.Printf("Inserting %d at index %d...\n", value, index)
			list.Insert(value, index)
		}
	}
	fmt.Printf("list size: %d\n", list.Size())
	fmt.Printf("list: %s\n", list.String())
	for i := 0; i < size; i++ {
		value, err := list.Get(i)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("list.Get(%d): %d\n", i, value)
	}
}
