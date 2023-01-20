package main

import (
	"fmt"
	"sync"
)

func main(){
	//go func(){
	//	for {
	//		//time.Sleep(time.Second*1)
	//		fmt.Println("range")
	//	}
	//}()
	f22()
	fmt.Println("exit")
}
func f22() {
	number_chan := make(chan bool)
	letter_chan := make(chan bool)
	wait := sync.WaitGroup{}
	go func(){
		i:=0
		for {
			select{
			case <-number_chan:
				i++
				fmt.Print(i)
				i++
				fmt.Print(i)
				letter_chan<-true
			}
		}
	}()
	wait.Add(1)

	go func(wait *sync.WaitGroup){
		i := 'A'-1
		for {
			select {
			case <-letter_chan:
				i++
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				if i == 'Z'{
					wait.Done()
					return
				}
				number_chan<-true
			}
		}
	}(&wait)
	number_chan<-true
	wait.Wait()
}