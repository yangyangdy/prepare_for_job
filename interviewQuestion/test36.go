package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string) //不带缓冲的channel，只有一个会执行成功其余的都会等待，会造成严重的内存泄露
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}
//判断一个channel是否close的方法
//val, ok := <-chan
func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}

