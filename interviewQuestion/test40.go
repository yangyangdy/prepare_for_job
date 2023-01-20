package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

/*
假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排列。
限时5秒，使用多个goroutine查找切片中是否存在给定值，在找到目标值或者超时后立刻结束所有goroutine的执行。


比如切片为：[23, 32, 78, 43, 76, 65, 345, 762, …… 915, 86]，查找的目标值为345，如果切片中存在目标值程序输出:"Found it!"
并且立即取消仍在执行查找任务的goroutine。
如果在超时时间未找到目标值程序输出:"Timeout! Not Found"，同时立即取消仍在执行查找任务的goroutine。
 */

//找到目标值或者超时后立刻结束所有goroutine的执行，需要用到计时器，channel和context
//用context.WithCancel创建一个上下文对象传递给每个执行任务的goroutine，
//外部在满足条件后（找到目标值或者已超时）通过调用上下文的取消函数来通知所有goroutine停止工作。
func main() {
	timer := time.NewTimer(time.Second * 5)
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan bool) //创建一个接收查找结果的通道的
	//...
	//for ... go SearchTarget
	select {
	case <-timer.C:
		fmt.Fprintln(os.Stderr, "Timeout! Not Found")
		cancel()
	case <-resultChan:
		fmt.Fprintf(os.Stdout,"Found it!\n")
		cancel()
	}
}

//执行任务的goroutine们如果找到目标值后需要通知外部等待任务执行的主goroutine，这个工作是典型的应用通道的场景，
//上面代码也已经看到了，我们创建了一个接收查找结果的通道，接下来要做的就是把它和上下文对象一起传递给执行任务的goroutine。
func SearchTarget(ctx context.Context, data []int, target int, resultChan chan bool) {
	for _, v := range data {
		//先判断是否时间到了，或者取消了
		select {
		//在执行查找任务的goroutine里接收上下文的取消信号，为了不阻塞查找任务，我们使用了select语句加default的组合
		case <- ctx.Done(): //ctx的done和cancel对应的
			fmt.Fprintf(os.Stdout, "Task cancelded! \n")
			return
		default:
		}
		//模拟一个耗时查找，这里只是对比值，真实开发中可以是其他操作
		fmt.Fprintf(os.Stdout, "v: %d \n", v)
		time.Sleep(time.Millisecond * 1500)
		if target == v {
			resultChan <- true
			return
		}
	}
}