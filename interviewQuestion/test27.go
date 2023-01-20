package main

import (
	"fmt"
	"sync"
	"time"
)
type singletone struct {

}
var once sync.Once
var instance *singletone
//单例模式
func main(){

}

func getInstance() *singletone {
	once.Do(func(){
		instance = &singletone{}
	})
	return instance
}

func getInstance1() *singletone {
	once.Do(func(){
		instance = &singletone{}
	})
	return instance
}

func f12() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch) //写channel的goroutine可能没有写完channel就被关闭了，这就会导致panic
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}