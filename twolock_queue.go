package queue

import (
	"sync"
)

// CQueue is a concurrent unbounded queue which uses two-Lock concurrent queue qlgorithm.
type CQueue[T any] struct {
	head  *cnode[T]
	tail  *cnode[T]
	hlock sync.Mutex
	tlock sync.Mutex
}

type cnode[T any] struct {
	value T
	next  *cnode[T]
}

// NewCQueue returns an empty CQueue.
func NewCQueue[T any]() *CQueue[T] {
	n := &cnode[T]{}
	return &CQueue[T]{head: n, tail: n}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *CQueue[T]) Enqueue(v T) {
	n := &cnode[T]{value: v}
	q.tlock.Lock()
	q.tail.next = n // Link node at the end of the linked list
	q.tail = n      // Swing Tail to node
	q.tlock.Unlock()
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *CQueue[T]) Dequeue() T {
	var t T
	q.hlock.Lock()
	n := q.head
	newHead := n.next
	if newHead == nil {
		q.hlock.Unlock()
		return t
	}
	v := newHead.value
	newHead.value = t
	q.head = newHead
	q.hlock.Unlock()
	return v
}
