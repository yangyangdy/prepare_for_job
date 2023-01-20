package my_obsatcle_question

//链表的归并排序
//涉及到1.两个排序链表的合并 2.找链表的中点 3.归并的模板

func mergeList1(L1 *ListNode, L2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for L1 != nil && L2 != nil {
		if L1.Val < L2.Val {
			cur.Next = L1
			L1 = L1.Next
		} else {
			cur.Next = L2
			L2 = L2.Next
		}
		cur = cur.Next
	}
	if L1 == nil {
		cur.Next = L2
	} else if L2 == nil {
		cur.Next = L1
	}
	return head.Next
}

func midList1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func sortList1(head *ListNode) *ListNode {
	//当只有一个元素的时候不用排
	if head == nil || head.Next == nil {
		return head
	}
	midNode := midList1(head)
	LA := head
	LB := midNode.Next
	midNode.Next = nil
	return mergeList1(sortList1(LA), sortList1(LB))
}