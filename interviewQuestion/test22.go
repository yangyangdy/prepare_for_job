package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// goroutine协程超时处理

func Do(ctx context.Context, wg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel() //
		wg.Done()
	}()
	done := make(chan struct{}, 1) //执行成功的channel
	go func(ctx context.Context) {
		fmt.Println("go gorutine")
		time.Sleep(time.Second * 6)
		done <- struct{}{} //发送完成的信号
	}(ctx)
	select {
	case <-ctx.Done(): //超时
		fmt.Printf("timeout,err:%v\n", ctx.Err())
	case <-time.After(3 * time.Second): //超时第二种方法
		fmt.Printf("after 1 sec.")
	case <-done: //程序正常结束
		fmt.Println("done")
	}
	fmt.Println("getout")
}
func main() {
	fmt.Println("main")
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	Do(ctx, &wg)
	wg.Wait()
	fmt.Println("finish")
}

var res int
var index int

func kthSmallest(root *TreeNode, k int) int {
	index = k
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		index--
		if index == 0 {
			res = root.Val
			return
		}
		dfs(root.Right)
	}
}

func dfs2(root *TreeNode) {
	if root != nil {
		dfs2(root.Left)
		index--
		if index == 0 {
			res = root.Val
			return
		}
		dfs(root.Right)
	}
}
