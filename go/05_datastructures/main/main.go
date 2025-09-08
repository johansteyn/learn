package main

import (
	"datastructures"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Datastructures")
	fmt.Println()

	size := 10
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Invalid size: %s\n", os.Args[1])
			os.Exit(1)
		}
		size = i
	}
	fmt.Printf("size: %d\n", size)
	rand.Seed(time.Now().UnixNano())
	start := time.Now().UnixMicro()

	fmt.Println("============ Linked List ============")
	list := datastructures.NewList()
	for i := 0; i < size; i++ {
		value := rand.Intn(size)
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
	for i := 0; i < size; i++ {
		//value := rand.Intn(size)
		value := i
		fmt.Printf("Searching for %d...", value)
		index, ok := list.Find(value)
		if ok {
			fmt.Printf(" found at index %d\n", index)
		} else {
			fmt.Printf(" not found\n")
		}
	}
	fmt.Printf("list: %s\n", list.String())
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
	fmt.Println()

	fmt.Println("============ Queue ============")
	queue := datastructures.NewQueue()
	for i := 0; i < size; i++ {
		value := rand.Intn(size)
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

	fmt.Println("============ Stack ============")
	stack := datastructures.NewStack()
	for i := 0; i < size; i++ {
		value := rand.Intn(size)
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

	fmt.Println("============ Binary Tree ============")
	tree := datastructures.NewTree()
	for i := 0; i < size; i++ {
		value := i
		//value := rand.Intn(size)
		//fmt.Printf("Adding %d...\n", value)
		tree.Add(value)
	}
	mid := time.Now().UnixMicro()
	//fmt.Printf("tree: %s\n", tree.String())
	tree.Print()
	fmt.Println()
	fmt.Printf("   Add time: %dµs\n", mid-start)
	size = tree.Size()
	depth := tree.Depth()
	mid = time.Now().UnixMicro()
	/*
		for i := 0; i < tree.Size(); i++ {
			value, err := tree.Get(i)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				break
			}
			fmt.Printf("tree.Get(%d): %d\n", i, value)
		}
	*/
	/*
		//maxDepth := 0
		for i := 0; tree.Size() > 0; i++ {
			value := i
			//value := rand.Intn(size)
			//_, ok := tree.Find(value)
			//if !ok {
			//	continue
			//}
			//fmt.Printf("Removing %d...\n", value)
			tree.Remove(value)
		}
	*/

	fmt.Println("============ Set ============")
	set1 := datastructures.NewSet[int]()
	for i := 0; i < size; i++ {
		n := rand.Intn(size)
		fmt.Printf("Adding %d to set1...\n", n)
		set1.Add(n)
	}
	var set2 *datastructures.Set[int]
	set2 = datastructures.NewSetFrom[int](0, 2, 4, 6, 8, 10)
	fmt.Printf("Size of set1: %d\n", set1.Size())
	fmt.Printf("Size of set2: %d\n", set2.Size())
	fmt.Printf("Contents of set1: %v\n", set1)
	fmt.Printf("Contents of set2: %v\n", set2)
	fmt.Println("Removing '10' from set2...")
	set2.Remove(10)
	fmt.Printf("Size of set2: %d\n", set2.Size())
	fmt.Printf("Contents of set2: %v\n", set2)
	fmt.Println("Iterating over set1:")
	seq := set1.Iter()
	for i := range seq {
		fmt.Printf(" %d\n", i)
	}
	fmt.Println("Iterating over set2:")
	seq = set2.Iter()
	for i := range seq {
		fmt.Printf(" %d\n", i)
	}
	printSet("Elements contained in set1", set1)
	printSet("Elements contained in set2", set2)
	printSet("Filtered elements in set1", set1.Filter(func(i int) bool { return i < 5 }))
	printSet("Doubled elements in set1", set1.Map(func(i int) int { return i * 2 }))
	printSet("Union of set1 and set2", set1.Union(set2))
	printSet("Intersection of set1 and set2", set1.Intersect(set2))
	printSet("Elements in set1 but not in set2", set1.Subtract(set2))
	printSet("Elements in set2 but not in set1", set2.Subtract(set1))
	fmt.Println()

	fmt.Println("Results...")
	end := time.Now().UnixMicro()
	fmt.Printf("Remove time: %dµs\n", end-mid)
	fmt.Printf("       Size: %d\n", size)
	fmt.Printf("      Depth: %d\n", depth)
	fmt.Printf("   Balances: %d\n", datastructures.Balances)
	fmt.Printf("  Rotations: %d\n", datastructures.Rotations)
	fmt.Println()
}

func printSet(label string, s *datastructures.Set[int]) {
	fmt.Printf("%s: ", label)
	count := 0
	for i := 0; ; i++ {
		if s.Contains(i) {
			fmt.Printf("%d ", i)
			count++
		}
		if count >= s.Size() {
			break
		}
	}
	fmt.Println()
}
