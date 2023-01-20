//package usual_err_go
package main

import (
	"errors"
	"fmt"
)
//在函数调用参数中，数组是值传递，无法通过修改数组类型的参数返回结果
func gotest1(){
	var a = []interface{}{1,2,3}
	fmt.Println(a)
	fmt.Println(a...)
}

//数组是值传递的

/*
golang中数组初始化的方式
1. var 数组变量名 [元素数量]T
var a [5]int 数组的长度也是数组类型的一部分
初始化的方式
1. var arr [3]int = [3]int{1,2,3}
2. var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
3. var arr = [3]int{1,2,3}
4. arr := [3]int{1,2,3}
5. var arr = [...]int{1,2,3} 让编译器自行推断数组的长度
6. arr := [5]int{0:1, 4:3} 0:1表示数组下标0对应的值为1
	fmt.Println(arr) //[1 0 0 0 3]
 */
func gotest2(){
	x := [3]int{1,2,3}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(x)
	}(x)

	fmt.Println(x)
}

//同名局部变量返回值被屏蔽
func gotest3()(err error){
	Bar := func()(error){
		return errors.New("ooop!")
	}
	if err := Bar();err != nil { //在局部作用域中命名的返回值会被同名的局部变量屏蔽
		return
	}
	return
}

func main(){
	e := gotest3()
	fmt.Println(e.Error())
}
