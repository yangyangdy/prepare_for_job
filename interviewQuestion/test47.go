package main

import "fmt"

type fpshandler interface {
	func1 (s string)
	func2 (s string)
}

type fpsTest struct {
	fpshandler
}

type fps1 struct {

}

func (*fps1)func1(s string) {
	fmt.Println("fps1 func1")
}
func (*fps1)func2(s string) {
	fmt.Println("fps1 func2")
}

type fps2 struct {

}

func (*fps2)func1(s string) {
	fmt.Println("fps2 func1")
}
func (*fps2)func2(s string) {
	fmt.Println("fps2 func2")
}

func main(){
	fps_test := fps2{}
	fps_test.func1("")
	fps_test.func2("")
}