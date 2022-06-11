# queue
```go
import "github.com/smallnest/queue"
```
Package queue provides multiple queue implementations. 
The lock-free and two-lock algorithms are from Michael and Scott. https://doi.org/10.1145/248052.248106

## Queue

A queue is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence.The operations of a queue make it a first-in-first-out (FIFO) data structure. In a FIFO data structure, the first element added to the queue will be the first one to be removed. 

Package queue defines `Queue` interface:

```go
type Queue[T any] interface {
	Enqueue(v T)
	Dequeue() T
}
```

Currently it contains three implementations:

- lock-free queue: `LKQueue`
- two-lock queue: `CQueue`
- slice-based queue: `SliceQueue`


## Benchmark


```sh
goos: darwin
goarch: amd64
pkg: github.com/smallnest/queue

BenchmarkQueue/lock-free_queue#4-8         	 4835329	       266.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkQueue/two-lock_queue#4-8          	 9112242	       168.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkQueue/slice-based_queue#4-8       	 8778811	       182.0 ns/op	      40 B/op	       0 allocs/op

BenchmarkQueue/two-lock_queue#32-8         	 9109314	       133.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkQueue/slice-based_queue#32-8      	 7939176	       171.8 ns/op	      54 B/op	       0 allocs/op
BenchmarkQueue/lock-free_queue#32-8        	 4735264	       253.8 ns/op	      16 B/op	       1 allocs/op

BenchmarkQueue/lock-free_queue#1024-8      	 4654297	       242.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkQueue/two-lock_queue#1024-8       	 7714422	       138.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkQueue/slice-based_queue#1024-8    	 8609463	       169.7 ns/op	      34 B/op	       0 allocs/op
```