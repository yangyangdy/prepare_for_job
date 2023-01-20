package main

import "fmt"

func main() {
	a := "hello"
	b := []byte(a)
	var c []byte
	c = b
	fmt.Println(&a)
	fmt.Println(&b[0])
	fmt.Println(&c[0])
}
