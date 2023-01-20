package main

import (
"fmt"
)
func main() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:] //->str1
	str2[1] = "new"
	fmt.Println(str1)
	//append会导致底层数组扩容生成新的数组，则不会影响到str1
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)

	//切片初始化 ,切片的[]中一定是空
	s := []int{1,2,3}
	//数组初始化，数组的[]中一定是有值的或者是"..."
	var r = [5]int{}
	var r1 = [...]int{}
}
