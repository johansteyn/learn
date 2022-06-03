package datastructures

import (
	"fmt"
)

type treeNode struct {
	value  int
	parent *treeNode
	left   *treeNode
	right  *treeNode
}

type BinaryTree struct {
	root *treeNode
	size int
}

func NewTree() BinaryTree {
	return BinaryTree{nil, 0}
}

func (t *BinaryTree) Size() int {
	return t.size
}

func (t *BinaryTree) Add(value int) {
	if t.size == 0 {
		newNode := &treeNode{value, nil, nil, nil}
		t.root = newNode
		t.size++
		return
	}
	t.root.add(value)
	t.size++
}

func (n *treeNode) add(value int) {
	if value <= n.value {
		if n.left == nil {
			newNode := &treeNode{value, n, nil, nil}
			n.left = newNode
		} else {
			n.left.add(value)
		}
	} else {
		if n.right == nil {
			newNode := &treeNode{value, n, nil, nil}
			n.right = newNode
		} else {
			n.right.add(value)
		}
	}
}

/*
func (t *BinaryTree) Insert(value int, index int) error {
	if index < 0 || index > t.size {
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
	t.size++
	return nil
}
*/

/*
func (l *LinkedList) Remove(index int) error {
	if index < 0 || index > l.size-1 {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
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
	if node == nil {
		panic("node not found in linked list")
	}
	if node.prev == nil {
		l.head = node.next
		if l.head != nil {
			l.head.prev = nil
		}
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		l.tail = node.prev
		if l.tail != nil {
			l.tail.next = nil
		}
	} else {
		node.next.prev = node.prev
	}
	l.size--
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
*/

func (t *BinaryTree) String() string {
	return fmt.Sprintf("%v", t.slice())
}

func (t *BinaryTree) slice() []int {
	s := make([]int, 0)
	/*
		node := t.root
		for i := 0; i < l.size; i++ {
			s[i] = node.value
			node = node.next
		}
		return s
	*/
	if t.root != nil {
		s = t.root.slice()
	}
	return s
}

func (n *treeNode) slice() []int {
	s := make([]int, 0)
	if n == nil {
		return s
	}
	s = append(s, n.value)
	leftSlice := n.left.slice()
	rightSlice := n.right.slice()
	s = append(leftSlice, s...)
	s = append(s, rightSlice...)
	return s
}
