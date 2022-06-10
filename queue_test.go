package queue

import "testing"

func TestQueue(t *testing.T) {
	queues := map[string]Queue[int]{
		"lock-free queue":   NewLKQueue[int](),
		"two-lock queue":    NewCQueue[int](),
		"slice-based queue": NewSliceQueue[int](0),
		"bounded queue":     NewBoundedQueue[int](100),
	}

	for name, q := range queues {
		t.Run(name, func(t *testing.T) {
			count := 100
			for i := 1; i <= count; i++ {
				q.Enqueue(i)
			}

			for i := 1; i <= count; i++ {
				v := q.Dequeue()
				if v == 0 {
					t.Fatalf("got a nil value")
				}
				if v != i {
					t.Fatalf("expect %d but got %v", i, v)
				}
			}
		})
	}

}
