package main

import (
	"fmt"
)

type m1 struct {
	a int
}
type m2 struct {
	b string
}
type model struct {
	mp map[int]*m1
	st []*m2
}

func main() {
	t1 := &model{}
	if t1.st == nil {
		fmt.Println("get 1")
	}
	w1 := make(map[int]*m1)
	w2 := make([]*m2, 2)
	t1.mp = w1
	t1.st = w2
	if t1.st == nil {
		fmt.Println("get 2")
	}
	//mp := map[int] string{
	//	1:"1",
	//	2:"2",
	//	3:"3",
	//}
	//if _, ok := mp[4]; !ok {
	//	fmt.Println("yes")
	//}else {
	//	fmt.Println("no")
	//}
	//s := make([]string, 2, 4)
	//example1(s, "hello", 10)
}

func example1(s []string, str string, i int) {
	panic("want stack trace")
}
