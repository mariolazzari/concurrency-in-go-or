package main

import (
	"fmt"
	"time"
)

var data int

func main() {
	// 	In Go, you can use the go keyword to run a function concurrently.
	// Doing so creates what’s called a goroutine
	go func() {
		data++
	}()

	// This is bad!
	time.Sleep(time.Second)

	if data == 0 {
		fmt.Println("Data is 0")
	}
}
