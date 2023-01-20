package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

//发布订阅者模型

type (
	subscriber chan interface{} //订阅者为一个管道，为啥用interface
	topicFunc func(v interface{}) bool //主题为一个过滤器
)

//发布者对象
type Publisher struct {
	m sync.RWMutex //读写锁
	buffer int //订阅队列的缓存大小
	timeout time.Duration //发布超时时间
	subscribers map[subscriber]topicFunc //订阅者信息，管道是可以作为key的
}
//构建一个发布者对象，可以设置发布超时时间和缓存队列的长度
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{ //为啥没有初始化锁
		buffer: buffer,
		timeout: publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

//添加一个新的订阅者，订阅全部的主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

//添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

//退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub) //删除map中的元素
	close(sub) //关闭channel
}

//发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock() //加读锁？
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

//关闭发布者对象，同时关闭所有的订阅者管道
func (p *Publisher) Close (){
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

//发送主题，可以容忍一定的超时
func (p *Publisher) sendTopic (sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup){
	defer wg.Done()
	if topic != nil && !topic(v) { //筛选器
		return
	}
	select {
	case sub <- v:
	case <- time.After(p.timeout):
	}
}

//在发布订阅模型中，每条消息都会传送给多个订阅者。发布者通常不会知道、也不关心哪一个订阅者正在接收主题消息。
//订阅者和发布者可以在运行时动态添加，是一种松散的耦合关系，这使得系统的复杂性可以随时间的推移而增长。
//在现实生活中，像天气预报之类的应用就可以应用这个并发模式。

func main() {

	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close() //关闭所有的订阅通道

	all := p.Subscribe() //返回的是订阅了全部主题的channel
	golang := p.SubscribeTopic(func(v interface{}) bool { //订阅筛选后的主题
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,  world!")
	p.Publish("hello, golang!")

	go func() {
		for  msg := range all {
			fmt.Println("all:", msg)
		}
	} ()

	go func() {
		for  msg := range golang {
			fmt.Println("golang:", msg)
		}
	} ()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}