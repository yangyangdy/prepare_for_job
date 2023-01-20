package main

import (
	"fmt"
	"time"
)

//实现一个简单的令牌桶限流
var fillInterval = time.Millisecond*100 //100ms触发一次
//往桶里面放令牌的时间间隔
var capacity = 100 // 每次放入桶中的令牌数
var tokenBucket = make(chan struct{}, capacity)

func main() {


	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		defer ticker.Stop()
		for {
			select { //select和channel共用
			case <-ticker.C: //当经过一个fillInterval后会触发
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}
	go fillToken()
	time.Sleep(time.Millisecond*11000)
}

//取令牌桶的操作, 每次只取一个令牌
func TakeAvailable(block bool) bool {
	var takenResult bool
	if block {
		select {
		case <-tokenBucket:
			takenResult = true
		}
	} else {
		select {
		case <-tokenBucket:
			takenResult = true
		default:
			takenResult = false
		}
	}
	return takenResult
}