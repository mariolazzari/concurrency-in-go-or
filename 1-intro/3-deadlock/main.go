package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func main() {
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		// Here we attempt to enter the critical section for the incoming value.
		v1.mu.Lock()
		// Here we use the defer statement to exit the critical section before printSum returns.
		defer v1.mu.Unlock()

		// Here we sleep for a period of time to simulate work (and trigger a deadlock)
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
