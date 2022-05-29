package datastructures

import (
	"fmt"
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

// Interesting...
// The "list" type is not exported, but it can be used indirectly in main.go
func New() list {
	return list{nil, nil, 0}
}

func (l *list) Size() int {
	return l.size
}

// A slow add function that traverses the entire list
func (l *list) Add1(value int) {
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
func (l *list) Add2(value int) {
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

func (l *list) Get(index int) (int, error) {
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

func (l *list) String() string {
	return fmt.Sprintf("%v", l.slice())
}

func (l *list) slice() []int {
	s := make([]int, l.size)
	node := l.head
	for i := 0; i < l.size; i++ {
		s[i] = node.value
		node = node.next
	}
	return s
}

// Global variable for a function that gets the temperature from a web service
// This implementation simply returns a hard-coded value
// It is replaced in the TestMain to "mock" the implementation
var GetTemperature = func() int {
	return 24
}
