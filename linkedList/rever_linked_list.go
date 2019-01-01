package linkedList

import "go.zhuzi.me/go/log"

/*
 * Definition for singly-linked list.
 * 反转单链表
 * 1->2->3->4->5->NULL
 * 5->4->3->2->1->NULL
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func Run() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5

	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	p := ReverseList(head)
	log.Error(p)
}

func ReverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil

	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}

	return pre
}

func OldReverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		tmp := pre // nil --- {1, nil}

		pre = cur // {1, 0xc42008a5} --- {2, 0xc42008a5}

		cur = cur.Next // {2, 0xc42008a5} --- {3, 0xc42008a5}

		pre.Next = tmp // nil => newHead = {1, nil} --- {1, nil} => newHead = {2, {1,nil}}....
	}
	return pre
}
