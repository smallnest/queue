package main

import (
	"sync"

	"github.com/smallnest/queue"
)

func main() {
	q := queue.NewCQueue[int]()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			q.Enqueue(1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			_ = q.Dequeue()
		}
	}()
	wg.Wait()
}
