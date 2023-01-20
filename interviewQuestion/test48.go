package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}
//如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的方法。
//反之不会

//当实现了一个接收者是值类型的方法，就可以自动生成一个接收者是对应指针类型的方法，因为两者都不会影响接收者。但是，当实现了一个接收者是指针类型的方法，如果此时自动生成一个接收者是值类型的方法，原本期望对接收者的改变（通过指针实现），现在无法实现，因为值类型会产生一个拷贝，不会真正影响调用者。

//值接收者
func (p Gopher) code() {
	p.language = "revise"
	fmt.Printf("I am coding %s language\n", p.language)
}

//指针接收者
func (p Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	var c coder = Gopher{"Go"}
	c.code()
	c.debug()
	fmt.Println(c)
}