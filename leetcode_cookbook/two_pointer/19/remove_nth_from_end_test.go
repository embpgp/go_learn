package remove_nth_from_end

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法一
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}

func printListNode(head *ListNode) {
	p := head
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println("")
}

func TestRemoveNthFromEnd(t *testing.T) {
	n1 := ListNode{1, nil}
	n2 := ListNode{2, nil}
	n3 := ListNode{3, nil}
	n4 := ListNode{4, nil}
	n5 := ListNode{5, nil}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	printListNode(&n1)
	removeNthFromEnd(&n1, 2)
	printListNode(&n1)

}
