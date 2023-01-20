package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	// var ch chan struct{}
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
	<-ch
}
