package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 5)
	data := []int{1, 2, 3, 10, 999, 8, 345, 7, 98, 33, 66, 77, 88, 68, 96}
	dataLen := len(data)
	size := 3
	target := 345
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan bool)
	for i:=0;i<dataLen;i+=size{
		end := i+size
		if end >= dataLen{
			end = dataLen-1
		}
		go SearchTarget1(ctx, data[i:end], target, resultChan)
	}
	select { //这里如果channel会被阻塞，知道超时或者找到
	case <-timer.C://计时器
		fmt.Println("timeout")
		cancel() //上下文用于通知其他有context的goroutine
	case <-resultChan:
		fmt.Println("found it")
		cancel() //这个会通知context取消
	}
	time.Sleep(time.Second*2)
}

func SearchTarget1(ctx context.Context, data []int, target int, resultChan chan bool) {
	for _, v := range data {
		select {
		case <- ctx.Done(): //用来判断是否继续执行
			fmt.Println("cancelded")
			return
		default:
		}
		if target == v{
			resultChan <- true
			return
		}
	}
}