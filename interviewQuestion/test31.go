package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a+b
	fmt.Println(index, a, b, ret)
	return ret
}

//defer在定义时会计算好调用函数的参数，所以会优先输出10， 20两个参数
//然后根据定义的顺序倒序执行
func main(){
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	s := make([]int, 5) //在初始化切片时指定了长度，所以追加数据时会从len(s)位置开始填充数据
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
