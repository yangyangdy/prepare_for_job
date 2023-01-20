package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ss := "/ssss"
	aaa := ss[1:]
	fmt.Println(aaa)
	c := 100
	var a *int = &c
	fmt.Println(a)
	*a = *a -10
	fmt.Println(a)
	fmt.Println(*a)
	return
	//创建一个子节点的context,3秒后自动超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	for {
		select {
		case <- ctx.Done():
			fmt.Println("timeout")
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(time.Second)
		}
	}
}

func f123(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	for {
		select {
		case <- ctx.Done():
			fmt.Println("timeout")
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(time.Second)
		}
	}
}
