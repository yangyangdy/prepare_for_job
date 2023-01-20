package my_obsatcle_question

type ListNode struct {
	Val int
	Next *ListNode
}
//链表的归并排序
func midList(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(L1 *ListNode, L2 *ListNode) *ListNode {
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
	if L1 == nil{
		cur.Next = L2
	} else if L2 == nil {
		cur.Next = L1
	}
	return head.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := midList(head)
	LA := head
	LB := mid.Next
	mid.Next = nil
	return mergeList(sortList(LA),sortList(LB))
}