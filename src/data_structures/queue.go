package data_structures

import (
	"errors"
	"fmt"
)

type Queue interface {
	Enqueue(n int)
	Dequeue() (int, error)
	IsEmpty() bool
}

type ArrQueue struct {
	head *node
	tail *node
	size int
}

type node struct {
	value int
	next  *node
}

func (q *ArrQueue) Enqueue(n int) {
	q.size += 1
	if q.head == nil {
		q.head = &node{n, nil}
		q.tail = q.head
		return
	}
	var tmp = &node{n, nil}
	q.tail.next = tmp
	q.tail = tmp
}

func (q *ArrQueue) Dequeue() (int, error) {
	if q.head == nil {
		return 0, errors.New("cannot dequeue empty queue")
	}
	var res = q.head.value
	q.head = q.head.next
	q.size -= 1
	return res, nil
}

func (q *ArrQueue) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}

func (q *ArrQueue) String() string {
	var el = q.head
	var arr = make([]int, 0, q.size)
	for el != nil {
		arr = append(arr, el.value)
		el = el.next
	}
	return fmt.Sprint(arr)
}
