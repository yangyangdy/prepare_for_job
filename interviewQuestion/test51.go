package main

import "fmt"

//素数筛

//返回生成自然数序列的管道2,3,4,...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
//然后是为每个素数构造一个筛子：将输入序列中是素数倍数的数提出，并返回新的序列，是一个新的管道
//管道过滤器：删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() //自然数序列: 2,3,4,...
	for i := 0; i < 100; i++ {
		prime := <-ch //新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) //基于新素数构造的过滤器
	}
}
