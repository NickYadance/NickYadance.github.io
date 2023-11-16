package codetop

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head, end *ListNode) (tail *ListNode) {
	if head != end {
		reverse(head.Next, end).Next = head
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 1; i < k && tail != nil; i = i + 1 {
		tail = tail.Next
	}
	if tail != nil {
		reverse(head, tail).Next = reverseKGroup(tail.Next, k)
		return tail
	} else {
		return head
	}
}

func Test25(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reversed := reverseKGroup(head, 2)
	fmt.Println(reversed.Val)
}
