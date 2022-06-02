package datastructures

import (
	"fmt"
)

type listNode struct {
	value int
	prev  *listNode
	next  *listNode
}

type LinkedList struct {
	head *listNode
	tail *listNode
	size int
}

func NewList() LinkedList {
	return LinkedList{nil, nil, 0}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Add(value int) {
	l.Insert(value, l.size)
}

func (l *LinkedList) Insert(value int, index int) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	newNode := &listNode{value, nil, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return nil
	}
	if index == 0 {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
		l.size++
		return nil
	}
	if index == l.size {
		// Appending
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		l.size++
		return nil
	}
	var node *listNode
	if index < l.size/2 {
		node = l.head
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		node = l.tail
		for i := l.size - 1; i > index; i-- {
			node = node.prev
		}
	}
	newNode.prev = node.prev
	if newNode.prev != nil {
		newNode.prev.next = newNode
	}
	newNode.next = node
	if newNode.next != nil {
		newNode.next.prev = newNode
	}
	l.size++
	return nil
}

func (l *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= l.size {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}
	node := l.head
	for i := 0; i < l.size; i++ {
		if index == i {
			return node.value, nil
		}
		node = node.next
	}
	return 0, fmt.Errorf("unexpected error")
}

func (l *LinkedList) String() string {
	return fmt.Sprintf("%v", l.slice())
}

func (l *LinkedList) slice() []int {
	s := make([]int, l.size)
	node := l.head
	for i := 0; i < l.size; i++ {
		s[i] = node.value
		node = node.next
	}
	return s
}
