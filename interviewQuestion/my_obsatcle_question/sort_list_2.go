package my_obsatcle_question

func midList2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList2(L1 *ListNode, L2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for L1 != nil && L2 != nil {
		if L1.Val <L2.Val {
			cur.Next = L1
			L1 = L1.Next
		} else {
			cur.Next = L2
			L2 = L2.Next
		}
		cur = cur.Next
	}
	if L1 != nil {
		cur.Next = L1
	} else if L2 != nil {
		cur.Next = L2
	}
	return head
}

func sortList2(head *ListNode) *ListNode {
	//出口
	if head == nil || head.Next == nil {
		return head
	}
	midNode := midList2(head)
	left := head
	right := midNode.Next
	midNode.Next = nil
	return mergeList2(sortList2(left), sortList2(right))
}

func main() {

}
