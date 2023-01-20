package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++ //竞态条件
			time.Sleep(time.Millisecond)
			count-- //竞态条件
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}
