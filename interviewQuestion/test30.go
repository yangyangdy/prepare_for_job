package main

import (
	"fmt"
	"runtime"
	"sync"
)
//闭包以引用的方式捕获外部变量

//输出结果决定来自于调度器优先调度哪个G,从runtime源码看，当创建一个G时，会优先放到下一个调度的runnext字段上作为下一次优先调度的G
//因此，最先输出的是最后创建的G
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i_: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i*: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

