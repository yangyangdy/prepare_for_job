package main

import "fmt"

//调用defer的流程，返回值赋值比defer表达式先执行
func ff1()(r int) {
	defer func(){
		r++
	}()
	return 0
}

func ff2()(r int) {
	t := 5
	defer func() {
		t =t+5
	}()
	return t
}

func ff3()(r int) {
	defer func(r int) {
		r = r+5
		fmt.Printf("value:%v\n", r)
	}(r)
	return 3
}

func main(){
	println(ff1())
	println(ff2())
	println(ff3())
}
