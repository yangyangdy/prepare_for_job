package main

import "fmt"
//求和定调用callback函数对结果进行处理
func sumWorker(data []int, callback func(int)) {
	sum := 0
	for _, num := range data {
		sum += num
	}
	callback(sum)
}

func Adder(x int) func(int) int{
	return func(y int) int{
		x += y
		fmt.Printf("x addr %p, x value %d\n", &x, x)
		return x
	}
}

//闭包的延迟绑定 ,最后输出的结果都会是16

//闭包保存/记录了它产生时的外部函数的所有环境。
//如同普通变量/函数的定义和实际赋值/调用或者说执行，是两个阶段。
//闭包也是一样，for-loop内部仅仅是声明了一个闭包，foo7()返回的也仅仅是一段闭包的函数定义，
//只有在外部执行了f7()时才真正执行了闭包，此时才闭包内部的变量才会进行赋值的操作。
//哎，如果这么说的话，岂不是应该抛出异常吗？因为val是一个比foo7()生命周期更短的变量啊？

func foo33(x int) []func() {
	var fs []func()
	values := []int{1,2,3,5}
	for _, val := range values {
		fs =append(fs, func(){
			fmt.Printf("foo3 val = %d\n", x+val)
		})
	}
	return fs //当返回的时候绑定最新的x+val的值
}

type testa struct{
	a *float64
}
//由函数及其相关引用环境组成的实体，捕获的本质就是引用传递而非值传递
func main() {
	ta := new(testa)
	fmt.Println(*ta.a)
	f3s := foo33(11)
	for _, f3 := range f3s {
		f3()
	}
	fmt.Println("----------------Adder()返回的匿名函数实例1----------------")
	closure := Adder(1)
	closure(100)
	closure(1000)
	closure(10000)

	fmt.Println("----------------Adder()返回的匿名函数实例2----------------")
	closure2 := Adder(10)
	closure2(1)
	closure2(1)
	closure2(1)
	closure2(1)

	return
}
