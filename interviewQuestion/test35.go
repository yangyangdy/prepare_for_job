package main

import "fmt"

const (
	a = iota
	b = iota
)

//iota每当遇到const时会重置为0
//iota在const组合声明中，每新增一行iota+
//跳值使用法 用"_"跳过某些值，用_来当垃圾桶占据某个值
//在const体重插入其他数值的赋值，计数不变
//在一个const组合声明中，如果没有指定表达式，则使用最近的空表达式，例如
/*
const(
a = iota*2 value 0
b = iota*3 value 3
c value 6(2*3)
d value 9(3*3)
)
 */
//表达式隐式使用法2
/*
const(
a,b = iota,iota+3 value 0,3
c,d value 1,4
)
 */
const (
	// name = "test1"
	c = iota
	d = iota
)
func main() {
	//ch := make(chan error)
	//fmt.Println("main1")
	//go func () {
	//	fmt.Println("go1")
	//	ch <- nil //阻塞
	//	fmt.Println("go2") //永远不会执行
	//}()

	fmt.Println(a) //value 0
	fmt.Println(b) //value 1
	fmt.Println(c) //value 1
	fmt.Println(d) //value 2
}
