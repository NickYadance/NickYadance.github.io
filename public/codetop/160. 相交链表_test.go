package codetop

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	i := headA
	j := headB
	for i != nil && j != nil {
		i = i.Next
		j = j.Next
	}
	for i != nil && headA != nil {
		i = i.Next
		headA = headA.Next
	}
	for j != nil && headB != nil {
		j = j.Next
		headB = headB.Next
	}
	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
