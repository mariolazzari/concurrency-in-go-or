# Concurrency in Go (o'Reilly)

## Introduction

### Race condition

A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.

```go

```

### Atomicity

Indivisible, or uninterruptible in a **context**

### Memory synchronization

```go
package main

import (
	"fmt"
	"sync"
)

// Here we add a variable that will allow our code to synchronize access to the data variable’s memory.
var memoryAccess sync.Mutex

var value int

func main() {

	go func() {
		// Here we declare that until we declare otherwise, our goroutine should have exclusive access to this memory.
		memoryAccess.Lock()
		value++
		// Here we declare that the goroutine is done with this memory.
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}
```

### Deadlock

A deadlocked program is one in which all concurrent processes are waiting on one another.

```go
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
```
