package main

import (
"fmt"
)

type People1 interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People1 = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}