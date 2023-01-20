package main

import (
	"fmt"
	"sync"
)

type ether1 struct {
	a int
	b string
	c float64
}

func main() {
	arr1 := make([]*ether1, 3)
	arr1[2] = &ether1{}

	arr := make([]*ether1, 0)
	test_arr := [...]int{1,2,3,4,5}
	waitGroup := sync.WaitGroup{}
	var lock sync.Mutex
	for i := range test_arr{
		waitGroup.Add(1)
		go func(val int){
			defer func(){
				waitGroup.Done()
			}()
			item := &ether1{
				a:val,
			}
			//slice不是并发安全的，需要加锁
			lock.Lock()
			defer lock.Unlock()
			arr = append(arr, item)
		}(test_arr[i])
	}
	waitGroup.Wait()
	for i := range arr{
		fmt.Println(arr[i].a)
		fmt.Println("********")
	}
}
