package main

import (
	"fmt"
	"time"
)

func addTwoNumbers(L1 *ListNode, L2 *ListNode) *ListNode {
	head := &ListNode{val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for L1 != nil || L2 != nil || carry != 0 {
		if L1 == nil {
			n1 = 0
		} else {
			n1 = L1.val
			L1 = L1.next
		}
		if L2 == nil {
			n2 = 0
		} else {
			n2 = L2.val
			L2 = L2.next
		}
		current.next = &ListNode{val: (n1 + n2 + carry) % 10}
		current = current.next
		carry = (n1 + n2 + carry) / 10
	}
	return head.next
}

func f_add_two(L1 *ListNode, L2 *ListNode) *ListNode {
	head := &ListNode{val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for L1 != nil || L2 != nil || carry != 0 {
		if L1 != nil {
			n1 = L1.val
			L1 = L1.next
		} else {
			n1 = 0
		}
		if L2 != nil {
			n2 = L2.val
			L2 = L2.next
		} else {
			n2 = 0
		}
		current.next = &ListNode{val: (n1 + n2 + carry) % 10}
		current = current.next
		carry = (n1 + n2 + carry) / 10
	}
	return head.next
}
func main() {

	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("hello1")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("hello2")
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func tsum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func tsum(nums []int, target int) []int {
	hastable := map[int]int{}
	for i, x := range nums {
		if p, ok := hastable[target-x]; ok {
			return []int{p, i}
		}
		hastable[x] = i
	}
	return nil
}

//最长的不重复字符的字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		if result < right-left+1 {
			result = right - left + 1
		}
	}
	return result
}
