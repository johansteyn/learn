package main

import (
	"datastructures"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Datastructures")
	fmt.Println()

	// Linked List
	list := datastructures.NewList()
	size := 10
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		value := rand.Intn(99) + 1
		if value%2 == 0 {
			fmt.Printf("Adding %2d to list...\n", value)
			list.Add(value)
		} else {
			index := 0
			if list.Size() > 0 {
				index = value % list.Size()
			}
			fmt.Printf("Inserting %2d to list at index %d...\n", value, index)
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
	fmt.Printf("Removing first node...\n")
	list.Remove(0)
	fmt.Printf("list size: %d\n", list.Size())
	fmt.Printf("list: %s\n", list.String())
	fmt.Printf("Removing middle node...\n")
	list.Remove(list.Size() / 2)
	fmt.Printf("list size: %d\n", list.Size())
	fmt.Printf("list: %s\n", list.String())
	fmt.Printf("Removing last node...\n")
	list.Remove(list.Size() - 1)
	fmt.Printf("list size: %d\n", list.Size())
	fmt.Printf("list: %s\n", list.String())

	// Queue
	queue := datastructures.NewQueue()
	for i := 0; i < size; i++ {
		value := rand.Intn(99) + 1
		fmt.Printf("Adding %2d to queue...\n", value)
		queue.Enqueue(value)
	}
	fmt.Printf("queue size: %d\n", queue.Size())
	fmt.Printf("queue: %s\n", queue.String())
	for queue.Size() > 0 {
		value, err := queue.Peek()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("Peeked:   %2d\n", value)
		value, err = queue.Dequeue()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("Dequeued: %2d\n", value)
	}
	fmt.Println()

	// Stack
	stack := datastructures.NewStack()
	for i := 0; i < size; i++ {
		value := rand.Intn(99) + 1
		fmt.Printf("Adding %2d to stack...\n", value)
		stack.Push(value)
	}
	fmt.Printf("stack size: %d\n", stack.Size())
	fmt.Printf("stack: %s\n", stack.String())
	for stack.Size() > 0 {
		value, err := stack.Peek()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("Peeked: %2d\n", value)
		value, err = stack.Pop()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("Popped: %2d\n", value)
	}
	fmt.Println()
}
