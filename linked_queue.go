package queue

// a non-threadsafe linked queue.
type LinkedQueue[T any] struct {
	head *cnode[T]
	tail *cnode[T]
	len  int
}

func (q *LinkedQueue[T]) Len() int {
	return q.len
}

func (q *LinkedQueue[T]) Reset() {
	n := &cnode[T]{}
	q.head = n
	q.tail = n
}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	n := &cnode[T]{}
	return &LinkedQueue[T]{head: n, tail: n}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *LinkedQueue[T]) Enqueue(v T) {
	n := &cnode[T]{value: v}
	q.tail.next = n // Link node at the end of the linked list
	q.tail = n      // Swing Tail to node
	q.len++
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *LinkedQueue[T]) Dequeue() T {
	var t T
	n := q.head
	newHead := n.next
	if newHead == nil {
		return t
	}
	v := newHead.value
	newHead.value = t
	q.head = newHead
	q.len--
	return v
}
