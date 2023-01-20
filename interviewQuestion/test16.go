package main

import "fmt"

func main(){
	var a []int
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	b:=make([]int, 5)
	copy(b[1:4], a)
	fmt.Println(b)
}
