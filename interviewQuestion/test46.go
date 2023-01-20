package main

import (
"fmt"
"os"
"runtime"
mytrace "runtime/trace"
	"sync"
	"sync/atomic"
"time"
)


//合理化内存分配的速度，提高赋值器的cpu利用率

var (
	stop_  int32
	count_ int64
	sum   time.Duration
)

func concat() { //创建800个 goroutine 执行字符串拼接
	wg := sync.WaitGroup{}
	for n := 0; n < 100; n++ {
		wg.Add(8) //每8个一批创建
		for i := 0; i < 8; i++ {
			go func() {
				s := "Go GC"
				s += " " + "Hello"
				s += " " + "World"
				_ = s
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	mytrace.Start(f)
	defer mytrace.Stop()

	go func() {
		var t time.Time
		for atomic.LoadInt32(&stop_) == 0 { //不断触发gc
			t = time.Now()
			runtime.GC() //进行gc操作
			sum += time.Since(t)
			count_++
		}
		fmt.Printf("GC spend avg: %v\n", time.Duration(int64(sum)/count_))
	}()

	concat()
	atomic.StoreInt32(&stop_, 1)
}
