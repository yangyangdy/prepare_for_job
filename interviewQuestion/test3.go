package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

// case1 case2 值传递和引用传递
// func 1
func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

// func 2
func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

// func 3
func foo3(){
	values := []int{1,2,3,4,5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}

// func 4
func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}
func foo4(){
	values := []int{1,2,3,5}
	for _, val := range values {
		go show(val)
	}
}

//func 5  func(){}() 加了括号表示要实例化即要执行
func foo5() {
	values := []int{1,2,3,5}
	for _, val := range values {
		go func(){
			fmt.Printf("foo5 val = %v\n", val)
		}()
	}
}

// func 6
var foo6Chan = make(chan int, 10)
func foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", val)
		}()
	}
}

// func 7
func foo7(x int) []func() {
	var fs []func()
	values := []int{1,2,3,5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

type name struct {
	firstname string
	lastname string
}

type pepole struct {
	myname *name
	sex    bool
}

func how() []int {
	s := []int{1, 2, 3}
	return s  //指针有可能发生逃逸
}

type SizeStruct struct {
	b byte
	in int64
}

func demoOne() {
	s := SizeStruct{'1', 23}
	fmt.Println("SizeStruct.b: ", unsafe.Sizeof(s.b),
		"SizeStruct.in: ", unsafe.Sizeof(s.in), "SizeStruct: ", unsafe.Sizeof(s))
	fmt.Println("SizeStruct Align: ", unsafe.Alignof(s))
}

//一个加法求和器 有着完整的词法环境， 闭包的性质
func adder() func(int) int {
	sum := 0
	return func(n int) int {
		sum += n
		return sum
	}
}

func EmptyInterface() {
	var r io.Reader
	fmt.Printf("io.Reader type[%v] value[%T]\n", r, r)
}

// 实现Reader
type space struct {
	str string
}

//实现io.Reader.Read
func (sp *space) Read(p []byte) (int, error){
	n := len(sp.str)
	for i := 0; i < n; i++ {
		p[i] = sp.str[i]
	}
	return n, nil
}

func InterfaceCom() {

	var i interface{} = 2
	t_i := reflect.TypeOf(i)
	v_i := reflect.ValueOf(i)
	fmt.Println(t_i, v_i)
	fmt.Println("=====InterfaceCom=====")
	sp := &space {
		str: "hello, pite\n hi sam",
	}

	reader := bufio.NewReader(sp)
	text, _ := reader.ReadString('\n') //读到换行
	text = strings.TrimSpace(text) //去掉前后的空格
	fmt.Printf("bufio.NewReader %#v\n", text)
}

//断言 等价接口类型的转换
//将接口类型变量的值(x)转成类型(T)，格式x.(T)
//T 可以是任意类型， T是实体类型，若想转换成功则T必须实现x的接口
//T 是接口类型，成功转换，则x的动态类型也必须实现了T接口，换言之T中所有方法都必须存在x接口内，或者x包含T
//nil接口的断言总是失败


func main(){
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w), reflect.ValueOf(w))
	f, ok := w.(*os.File)
	if ok {
		t_f := reflect.TypeOf(f)
		v_f := reflect.ValueOf(f)
		fmt.Println(t_f, &v_f)
	}
	//InterfaceCom()
	//EmptyInterface()
	//a := adder()
	//fmt.Println(a(0), a(2), a(4))
	//匿名函数
	//sq := func (x int) int {
	//	return x * x
	//}
	//fmt.Println("anonymous sq1:=", sq(2))
	//demoOne()
	//_ = how() //如果没有接收不会发生逃逸
	//fmt.Println("escape testing")
}
