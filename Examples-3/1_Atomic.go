package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
)

func main() {
	wg.Add(4) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Java")
	go increment("Golang")
	go atomicFunc()

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

func increment(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		atomic.AddInt32(&counter, 1)
		//counter++
		//runtime.Gosched() // Yield the thread and be placed back in queue.
	}
}

func atomicFunc() {
	wg.Done()
	counter += 30
}
