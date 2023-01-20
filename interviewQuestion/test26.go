package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverString(s string)(string, bool) {
	str := []rune(s) //遍历字符串之前需要先将其转换为rune形式！！!
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i:=0;i<l/2;i++{
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str), true
}
func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}



type People struct {
	Name string
}

//String实现了Stringer的接口
//会导致循环依赖？为什么
func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	p.String()
}
