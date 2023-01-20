package main

import (
	"fmt"
	"sync"
)

//周期打印字母和数字
func pringWordAndNum() {
	letter_chan := make(chan bool, 1)
	number_chan := make(chan bool, 1) //用于控制同步
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func(){
		defer func(){
			fmt.Println("exit...")
			wait.Done()
		}()
		i:=0
		j:='A'
		for {
			select{
			case <-number_chan:
				fmt.Println(i)
				i++
				letter_chan<-true //如果chan是阻塞式的则会阻塞在这里
			case <-letter_chan:
				fmt.Println(string(j))
				j++
				if j == 'Z'+1 {
					//wait.Done()
					return
				}
				number_chan<-true
			}
		}
	}()
	number_chan<-true
	wait.Wait()
	//time.Sleep(1*time.Second)
	return
}

func main(){
	pringWordAndNum()
}