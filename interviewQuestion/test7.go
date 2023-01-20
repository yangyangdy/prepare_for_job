package main

import "fmt"

type myinter interface{
	func1(name string)error
	func2()
}
func main() {
	var c int = 3
	c = c << 2
	fmt.Println(c)
	m := make(map[string]string)
	m["helloworld"] = "yes"
	fmt.Printf("m cache addr: %p\n", &m)
	hello(m)
	fmt.Printf("%v\n", m)

	s := []string{"aa", "bb", "cc"}
	fmt.Printf("%p\n", s)
	hellos(s)
	fmt.Println(s)
}

func hello(p map[string]string) {
	fmt.Printf("p cache addr: %p\n", &p)
	p["helloworld"] = "no"
}

func hellos(s []string) {
	fmt.Printf("%p\n", s)
	s[0]="dd"
}
