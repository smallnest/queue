package queue

// Queue is a FIFO data structure.
// Enqueue puts a value into its tail,
// Dequeue removes a value from its head.
type Queue[T any] interface {
	Enqueue(v T)
	Dequeue() T
	Len() int
	Reset()
}
