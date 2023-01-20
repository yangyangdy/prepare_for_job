package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//stu会被复用,每次循环会将集合中的值复制给这个变量，因此，会导致最后m中的map中存储的都是stus最后一个student的值
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Println(stu)
	}
	for _, item := range m{
		fmt.Println(*item)
	}

}

func main(){
	pase_student()
}
