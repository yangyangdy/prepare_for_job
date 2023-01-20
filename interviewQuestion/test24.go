package main

import (
	"fmt"
	"sync"
	//"time"
)

//周期打印数字和字母
//通过两个无容量的chan来控制同步
//利用waitGroup来控制协程返回后主协程才退出


func main(){
	f11()
	//fmt.Println(string('Z'+1))
	//letter, number := make(chan bool), make(chan bool)
	// wait := sync.WaitGroup{}
	//
	// go func() {
	//	i:=1
	//	for {
	//		select {
	//		case <-number:
	//			fmt.Print(i)
	//			i++
	//			fmt.Print(i)
	//			i++
	//			time.Sleep(1*time.Second)
	//			letter <- true //往letter的chan里放入true
	//		}
	//	}
	// }()
	// wait.Add(1)
	// go func(wait *sync.WaitGroup) {
	//	 i:='A'
	//	 for {
	//		 select {
	//		 case <- letter:
	//			 if i >= 'Z'{
	//				 fmt.Println()
	//				 wait.Done()
	//				 return
	//			 }
	//			 fmt.Print(string(i))
	//			 i++
	//			 fmt.Print(string(i))
	//			 i++
	//			 time.Sleep(1*time.Second)
	//			 number <- true
	//		 }
	//	 }
	// }(&wait)
	// number<-true //启动
	// wait.Wait()  //等待协程结束
}

func f11() {
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
		i:=int('A')-1
		for {
			select {
			case <-letter_chan:
				i++
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				if i == int('Z'){
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

