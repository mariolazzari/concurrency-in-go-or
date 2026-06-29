package main

import (
	"fmt"
	"time"
)

var data int

func main() {
	go func() {
		data++
	}()

	// This is bad!
	time.Sleep(time.Second)

	if data == 0 {
		fmt.Println("Data is 0")
	}
}
