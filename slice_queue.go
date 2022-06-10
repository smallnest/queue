package queue

import "sync"

// SliceQueue is an unbounded queue which uses a slice as underlying.
type SliceQueue[T any] struct {
	data []T
	mu   sync.Mutex
}

// NewSliceQueue returns an empty queue.
func NewSliceQueue[T any](n int) (q *SliceQueue[T]) {
	return &SliceQueue[T]{data: make([]T, 0, n)}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *SliceQueue[T]) Enqueue(v T) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *SliceQueue[T]) Dequeue() T {
	var t T
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return t
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
