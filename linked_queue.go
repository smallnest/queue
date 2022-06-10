package queue

// a non-threadsafe linked queue.
type linkedQueue[T any] struct {
	head *cnode[T]
	tail *cnode[T]
}

func newlinkedQueue[T any]() *linkedQueue[T] {
	n := &cnode[T]{}
	return &linkedQueue[T]{head: n, tail: n}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *linkedQueue[T]) Enqueue(v T) {
	n := &cnode[T]{value: v}
	q.tail.next = n // Link node at the end of the linked list
	q.tail = n      // Swing Tail to node
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *linkedQueue[T]) Dequeue() T {
	var t T
	n := q.head
	newHead := n.next
	if newHead == nil {
		return t
	}
	v := newHead.value
	newHead.value = t
	q.head = newHead
	return v
}
