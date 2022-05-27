package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type node struct {
	value int
	next  *node
}

type list struct {
	head *node
	size int
}

func New() list {
	return list{nil, 0}
}

func (l *list) add(value int) {
	l.size++
	n := &node{value, nil}
	if l.head == nil {
		l.head = n
		return
	}
	for tail := l.head; tail != nil; tail = tail.next {
		if tail.next == nil {
			tail.next = n
			return
		}
	}
}

func (l *list) get(index int) (int, error) {
	if index < 0 || index >= l.size {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}
	//for node := l.head, i := 0; node != nil; node = node.next, i++ {
	node := l.head
	for i := 0; i < l.size; i++ {
		if index == i {
			return node.value, nil
		}
		node = node.next
	}
	return 0, fmt.Errorf("unexpected error")
}

// TODO: A get method that returns the value at the specified index, returning an error if not found

func (l *list) String() string {
	s := ""
	for node := l.head; node != nil; node = node.next {
		s = s + " " + strconv.Itoa(node.value)
	}
	return s
}

func main() {
	fmt.Println("Linked List")
	fmt.Println()

	list := New()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		list.add(rand.Intn(100))
	}
	fmt.Printf("list: %s\n", list.String())
	fmt.Printf("list.size: %d\n", list.size)
	// Deliberately going out of bounds here - to be covered in a test...
	for i := 0; i < 11; i++ {
		value, err := list.get(i)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("list.get(%d): %d\n", i, value)

	}
}
