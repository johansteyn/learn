package datastructures

import (
	"fmt"
)

type Queue struct {
	list LinkedList
}

func NewQueue() Queue {
	return Queue{NewList()}
}

// Q: Should it not be possible to call the linked list's Size method directly?
func (q *Queue) Size() int {
	return q.list.size
}

func (q *Queue) Enqueue(value int) {
	q.list.Add(value)
}

func (q *Queue) Dequeue() (int, error) {
	if q.list.Size() <= 0 {
		return 0, fmt.Errorf("empty queue")
	}
	value, err := q.list.Get(0)
	if err != nil {
		return 0, err
	}
	err = q.list.Remove(0)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (q *Queue) Peek() (int, error) {
	if q.list.Size() <= 0 {
		return 0, fmt.Errorf("empty queue")
	}
	return q.list.Get(0)
}

func (q *Queue) String() string {
	return q.list.String()
}
