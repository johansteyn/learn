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

func (t *BinaryTree) Remove(value int) error {
	if t.size == 0 {
		return fmt.Errorf("value not found")
	}
	if t.size == 1 && t.root.value != value {
		return fmt.Errorf("value not found")
	}
	newRoot, err := t.root.remove(value)
	if err != nil {
		return err
	}
	if newRoot != t.root {
		t.root = newRoot
	}
	t.size--
	return nil
}

func (n *treeNode) remove(value int) (*treeNode, error) {
	if value < n.value {
		if n.left == nil {
			return nil, fmt.Errorf("value not found")
		}
		return n.left.remove(value)
	} else if value > n.value {
		if n.right == nil {
			return nil, fmt.Errorf("value not found")
		}
		return n.right.remove(value)
	}
	return n.delete(), nil
}

// TODO: adjust weights...
func (n *treeNode) delete() *treeNode {
	var replacement *treeNode
	if n.left == nil {
		replacement = n.right
	} else if n.right == nil {
		replacement = n.left
	} else {
		// Take the right child, and hang the left child off its leftmost child
		replacement = n.right
		leftmost := replacement
		for node := replacement.left; node != nil; node = node.left {
			leftmost = node
		}
		leftmost.left = n.left
		leftmost.left.parent = leftmost
	}
	if replacement != nil {
		replacement.parent = n.parent
	}
	if n.parent != nil {
		if n == n.parent.left {
			n.parent.left = replacement
		} else {
			n.parent.right = replacement
		}
	}
	root := replacement
	if root == nil {
		root = n.parent
	}
	if root != nil {
		for node := root.parent; node != nil; node = node.parent {
			root = node
		}
	}
	return root
}

// Does an in-order traversal (not used anywhere)
func (n *treeNode) traverse() {
	if n.left != nil {
		n.left.traverse()
	}
	fmt.Printf("Visited node: %d\n", n.value)
	if n.right != nil {
		n.right.traverse()
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
