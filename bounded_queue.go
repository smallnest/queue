package queue

import (
	"sync"
)

// BoundedQueue is a threadsafe bounded queue.
type BoundedQueue[T any] struct {
	capacity int
	q        Queue[T]
	cond     *sync.Cond
}

// NewBoundedQueue create a BoundedQueue.
func NewBoundedQueue[T any](n int, q Queue[T]) *BoundedQueue[T] {
	return &BoundedQueue[T]{
		capacity: n,
		q:        q,
		cond:     sync.NewCond(&sync.Mutex{}),
	}
}

// Enqueue puts the given value v at the tail of the queue.
// If this queue if full, the caller will be blocked.
func (q *BoundedQueue[T]) Enqueue(v T) {
	q.cond.L.Lock()
	for q.q.Len() == q.capacity {
		q.cond.Wait()
	}
	q.q.Enqueue(v)
	q.cond.L.Unlock()

	// change the condition
	q.cond.Broadcast()
}

// Dequeue removes and returns the value at the head of the queue.
// It will be blocked if the queue is empty.
func (q *BoundedQueue[T]) Dequeue() T {
	q.cond.L.Lock()
	for q.Len() == 0 {
		q.cond.Wait()
	}
	v := q.q.Dequeue()
	q.cond.L.Unlock()

	// change the condition
	q.cond.Broadcast()

	return v
}

// Len returns length of this queue.
func (q *BoundedQueue[T]) Len() int {
	return q.q.Len()
}

func (q *BoundedQueue[T]) Reset() {
	q.cond.L.Lock()
	q.q.Reset()
	q.cond.L.Unlock()
	q.cond.Broadcast()
}
