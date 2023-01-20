package main

import (
	"fmt"
	"sync"
)

var hits struct {
	sync.Mutex
	n int
}

func main() {
	var wg sync.WaitGroup
	N := 10
	// 启动100个goroutine对匿名结构体的成员n同时进行读写操作
	wg.Add(N)
	for i:=0; i<100; i++ {
		go func() {
			defer wg.Done()
			hits.Lock()
			defer hits.Unlock()
			hits.n++
		}()
	}
	wg.Wait()
	fmt.Println(hits.n)
}