package queue

// a non-threadsafe linked queue.
type linkedQueue struct {
	head *cnode
	tail *cnode
}

func newlinkedQueue() *linkedQueue {
	n := &cnode{}
	return &linkedQueue{head: n, tail: n}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *linkedQueue) Enqueue(v interface{}) {
	n := &cnode{value: v}
	q.tail.next = n // Link node at the end of the linked list
	q.tail = n      // Swing Tail to node
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *linkedQueue) Dequeue() interface{} {
	n := q.head
	newHead := n.next
	if newHead == nil {
		return nil
	}
	v := newHead.value
	newHead.value = nil
	q.head = newHead
	return v
}
