package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//waitgroup 防止主协程退出
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// <-chan表示只能接收
		// chan<-表示只能发送
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}
	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second*10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan bool, 1)

	go time.AfterFunc(timeout, func(){
		ch <- true
	})

	go func(){
		wg.Wait()
		ch <- false
	}()

	return <-ch
}