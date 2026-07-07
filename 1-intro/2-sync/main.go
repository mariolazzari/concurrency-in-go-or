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
