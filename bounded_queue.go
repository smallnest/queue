package queue

import (
	"sync"
	"sync/atomic"
)

// BoundedQueue is a threadsafe bounded queue.
type BoundedQueue[T any] struct {
	capacity uint32
	len      uint32
	q        *linkedQueue[T]
	cond     *sync.Cond
}

// NewBoundedQueue create a BoundedQueue.
func NewBoundedQueue[T any](n uint32) *BoundedQueue[T] {
	return &BoundedQueue[T]{
		capacity: n,
		q:        newlinkedQueue[T](),
		cond:     sync.NewCond(&sync.Mutex{}),
	}
}

// Enqueue puts the given value v at the tail of the queue.
// If this queue if full, the caller will be blocked.
func (q *BoundedQueue[T]) Enqueue(v T) {
	q.cond.L.Lock()
	for q.len == q.capacity {
		q.cond.Wait()
	}
	q.q.Enqueue(v)
	q.cond.L.Unlock()

	// change the condition
	atomic.AddUint32(&q.len, 1)
	q.cond.Broadcast()
}

// Dequeue removes and returns the value at the head of the queue.
// It will be blocked if the queue is empty.
func (q *BoundedQueue[T]) Dequeue() T {
	q.cond.L.Lock()
	for atomic.LoadUint32(&q.len) == 0 {
		q.cond.Wait()
	}
	v := q.q.Dequeue()
	q.cond.L.Unlock()

	// change the condition
	atomic.AddUint32(&q.len, ^uint32(0))
	q.cond.Broadcast()

	return v
}

// Len returns length of this queue.
func (q *BoundedQueue[T]) Len() int {
	l := atomic.LoadUint32(&q.len)
	return int(l)
}
