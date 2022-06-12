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
	newRoot := t.root.add(value)
	if newRoot.parent == nil {
		t.root = newRoot
	}
	t.size++
}

func (n *treeNode) add(value int) *treeNode {
	if value <= n.value {
		if n.left == nil {
			n.left = &treeNode{value, n, nil, nil, 1}
		} else {
			n.left.add(value)
		}
	} else {
		if n.right == nil {
			n.right = &treeNode{value, n, nil, nil, 1}
		} else {
			n.right.add(value)
		}
	}
	return n.rebalance()
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
	//fmt.Printf("*** %d.remove(%d)\n", n.value, value)
	if value < n.value {
		if n.left == nil {
			return n, fmt.Errorf("value not found")
		}
		return n.left.remove(value)
	} else if value > n.value {
		if n.right == nil {
			return n, fmt.Errorf("value not found")
		}
		return n.right.remove(value)
	}
	return n.delete(), nil
}

func (n *treeNode) delete() *treeNode {
	//fmt.Printf("*** %d.delete()\n", n.value)
	var replacement *treeNode
	if n.left == nil {
		replacement = n.right
	} else if n.right == nil {
		replacement = n.left
	} else {
		if n.left.weight > n.right.weight {
			for n.left.right != nil {
				n.left.srl()
				n.left.left.rebalance()
			}
			replacement = n.left
			replacement.right = n.right
			replacement.right.parent = replacement
		} else {
			for n.right.left != nil {
				n.right.srr()
				n.right.right.rebalance()
			}
			replacement = n.right
			replacement.left = n.left
			replacement.left.parent = replacement
		}
	}
	if replacement != nil {
		replacement = replacement.rebalance()
		replacement.parent = n.parent
	}
	if n.parent != nil {
		if n == n.parent.left {
			n.parent.left = replacement
		} else {
			n.parent.right = replacement
		}
	}
	if replacement == nil {
		replacement = n.parent
	}
	for node := replacement; node != nil; node = node.parent {
		node = node.rebalance()
		if node != nil {
			replacement = node
		}
	}
	return replacement
}

func (n *treeNode) root() *treeNode {
	if n == nil {
		return nil
	}
	root := n
	for ; n.parent != nil; n = n.parent {
		root = n.parent
	}
	return root
}

const tolerance = 25

var Balances int = 0
var Rotations int = 0

func (n *treeNode) rebalance() *treeNode {
	if n.left != nil {
		n.left.adjustWeight()
	}
	if n.right != nil {
		n.right.adjustWeight()
	}
	n.adjustWeight()
	node := n
	for ; !node.isBalanced(); node = node.balance() {
	}
	return node
}

func (n *treeNode) balance() *treeNode {
	//fmt.Printf("*** %d.balance()\n", n.value)
	Balances++
	if n.left == nil && n.right == nil {
		// Do nothing, but drop through to adjust weights...
	} else if n.left == nil {
		if n.right.weight > 1 {
			if n.right.left != nil {
				// Double rotation
				n.right.srr()
				n.right.right.rebalance()
			}
			return n.srl()
		}
	} else if n.right == nil {
		if n.left.weight > 1 {
			if n.left.right != nil {
				// Double rotation
				n.left.srl()
				n.left.left.rebalance()
			}
			return n.srr()
		}
	} else {
		percent := 100 * n.left.weight / (n.left.weight + n.right.weight)
		if percent < tolerance {
			if n.right.left != nil {
				if n.right.right == nil || n.right.left.weight > n.right.right.weight {
					// Double rotation
					n.right.srr()
					n.right.right.rebalance()
				}
			}
			return n.srl()
		} else if percent > 100-tolerance {
			if n.left.right != nil {
				if n.left.left == nil || n.left.right.weight > n.left.left.weight {
					// Double rotation
					n.left.srl()
					n.left.left.rebalance()
				}
			}
			return n.srr()
		}
	}
	if n.left != nil {
		n.left.adjustWeight()
	}
	if n.right != nil {
		n.right.adjustWeight()
	}
	n.adjustWeight()
	return n
}

func (n *treeNode) isBalanced() bool {
	if n == nil {
		return true
	}
	if n.left == nil && n.right == nil {
		return true
	}
	if n.left == nil && n.right.weight > 1 {
		return false
	}
	if n.right == nil && n.left.weight > 1 {
		return false
	}
	if n.left != nil && n.right != nil {
		percent := 100 * n.left.weight / (n.left.weight + n.right.weight)
		if percent < tolerance || percent > 100-tolerance {
			return false
		}
	}
	return true
}

func (n *treeNode) srl() *treeNode {
	//fmt.Printf("*** %d.srl()\n", n.value)
	Rotations++
	n.right.parent = n.parent
	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = n.right
		} else {
			n.parent.right = n.right
		}
	}
	n.parent = n.right
	n.right = n.right.left
	if n.right != nil {
		n.right.parent = n
	}
	n.parent.left = n
	n.adjustWeight()
	if n.parent != nil {
		n.parent.adjustWeight()
	}
	return n.parent
}

func (n *treeNode) srr() *treeNode {
	//fmt.Printf("*** %d.srr()\n", n.value)
	Rotations++
	n.left.parent = n.parent
	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = n.left
		} else {
			n.parent.right = n.left
		}
	}
	n.parent = n.left
	n.left = n.left.right
	if n.left != nil {
		n.left.parent = n
	}
	n.parent.right = n
	n.adjustWeight()
	if n.parent != nil {
		n.parent.adjustWeight()
	}
	return n.parent
}

func (n *treeNode) adjustWeight() {
	n.weight = 1
	if n.left != nil {
		n.weight += n.left.weight
	}
	if n.right != nil {
		n.weight += n.right.weight
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

func (t *BinaryTree) Print() {
	if t.root == nil {
		fmt.Printf("◁\n")
		return
	}
	if t.root.right != nil {
		t.root.right.print(true, " ")
	}
	fmt.Printf("◁%d(%d)\n", t.root.value, t.root.weight)
	if t.root.left != nil {
		t.root.left.print(false, " ")
	}
}

func (n *treeNode) print(isRight bool, indent string) {
	if n.right != nil {
		addIndent := "┃ "
		if isRight {
			addIndent = "  "
		}
		n.right.print(true, indent+addIndent)
	}
	fmt.Print(indent)
	if isRight {
		fmt.Print("┏")
	} else {
		fmt.Print("┗")
	}
	fmt.Print("━")
	fmt.Printf("%d(%d)\n", n.value, n.weight)
	if n.left != nil {
		addIndent := "  "
		if isRight {
			addIndent = "┃ "
		}
		n.left.print(false, indent+addIndent)
	}
}

func (t *BinaryTree) Depth() int {
	if t.size <= 2 {
		return t.size
	}
	return t.root.depth()
}

func (n *treeNode) depth() int {
	depth := 1
	if n.left != nil {
		depth = 1 + n.left.depth()
	}
	if n.right != nil {
		d := 1 + n.right.depth()
		if d > depth {
			depth = d
		}
	}
	return depth
}
