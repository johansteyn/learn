package datastructures

import (
	"fmt"
)

type treeNode struct {
	value  int
	parent *treeNode
	left   *treeNode
	right  *treeNode
	weight int
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
		newNode := &treeNode{value, nil, nil, nil, 1}
		t.root = newNode
		t.size++
		return
	}
	t.root.add(value)
	t.size++
}

// TODO: Balance tree based on weights...
func (n *treeNode) add(value int) {
	if value <= n.value {
		if n.left == nil {
			newNode := &treeNode{value, n, nil, nil, 1}
			n.left = newNode
		} else {
			n.left.add(value)
		}
	} else {
		if n.right == nil {
			newNode := &treeNode{value, n, nil, nil, 1}
			n.right = newNode
		} else {
			n.right.add(value)
		}
	}
	n.weight++
}

func (t *BinaryTree) Remove(index int) error {
	if index < 0 || index > t.size-1 {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	// TODO...
	t.size--
	return nil
}

func (t *BinaryTree) Get(index int) (int, error) {
	if index < 0 || index >= t.size {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}
	return t.root.get(index)
}

func (n *treeNode) get(index int) (int, error) {
	leftWeight := 0
	if n.left != nil {
		leftWeight = n.left.weight
	}
	if index == leftWeight {
		return n.value, nil
	} else if index < leftWeight {
		return n.left.get(index)
	} else if index > leftWeight {
		return n.right.get(index - leftWeight - 1)
	}
	return -1, nil
}

// Does an in-order traversal
func (n *treeNode) traverse() {
	if n.left != nil {
		n.left.traverse()
	}
	fmt.Printf("Visited node: %d\n", n.value)
	if n.right != nil {
		n.right.traverse()
	}
}

func (t *BinaryTree) Find(value int) (index int, ok bool) {
	if t.root == nil {
		return 0, false
	}
	x := 0
	if t.root.left != nil {
		x = t.root.left.weight
	}
	return t.root.find(value, x)
}

func (n *treeNode) find(value int, x int) (index int, ok bool) {
	if value == n.value {
		return x, true
	}
	if value < n.value {
		if n.left == nil {
			return 0, false
		}
		y := x - 1
		if n.left.right != nil {
			y -= n.left.right.weight
		}
		return n.left.find(value, y)
	} else {
		if n.right == nil {
			return 0, false
		}
		y := x + 1
		if n.right.left != nil {
			y += n.right.left.weight
		}
		return n.right.find(value, y)
	}
}

func (t *BinaryTree) String() string {
	return fmt.Sprintf("%v", t.slice())
}

func (t *BinaryTree) slice() []int {
	s := make([]int, 0)
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
