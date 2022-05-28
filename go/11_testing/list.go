package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type node struct {
	value int
	next  *node
}

type list struct {
	head *node
	tail *node
	size int
}

func New() list {
	return list{nil, nil, 0}
}

// A slow add function that travereses the entire list
func (l *list) add1(value int) {
	l.size++
	newNode := &node{value, nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	for n := l.head; n != nil; n = n.next {
		if n.next == nil {
			n.next = newNode
			return
		}
	}
}

// A faster add function that uses the tail node
func (l *list) add2(value int) {
	l.size++
	newNode := &node{value, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	l.tail.next = newNode
	l.tail = newNode
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
	size := getTemperature()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		list.add1(rand.Intn(100))
	}
	fmt.Printf("list: %s\n", list.String())
	fmt.Printf("list.size: %d\n", list.size)
	// Deliberately going out of bounds here
	for i := 0; i < size+1; i++ {
		//for i := 0; i < size; i++ {
		value, err := list.get(i)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("list.get(%d): %d\n", i, value)
	}
}

// Global variable for a function that gets the temperature from a web service
// This implementation simply returns a hard-coded value
// It is replaced in the TestMain to "mock" the implementation
var getTemperature = func() int {
	return 10
}
