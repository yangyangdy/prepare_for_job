package main

import (
	"fmt"
	"time"
)

// chan<- 表示数据进入通道，把数据写入通道，对于调用者而言就是发送
// <-chan 表示数据从通道出来，调用者得到通道数据，就是接收

//递归的方式进行链表排序
type ListNode struct {
	Val int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	mid := midList(head)
	LA := head
	LB := mid.next
	mid.next = nil
	return merge(SortList(LA), SortList(LB))
}

func merge(L1 *ListNode, L2 *ListNode) *ListNode {
	var res = &ListNode{}
	next := res
	curA := L1
	curB := L2
	for curA != nil && curB != nil {
		if curA.val > curB.val {
			next.next = &ListNode{}
		}
	}
}

func midList(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
func main() {
	c := make(chan int, 10) // 不使用带缓冲区的channel
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c) // 关闭chan
}

// 只能向chan里send数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send readey ", i)
		c <- i
		fmt.Println("send ", i)
	}
}

// 只能接收chan中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received ", i)
	}
}

