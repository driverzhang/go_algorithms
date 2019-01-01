package main

import (
	"gitHub/go_algorithms/linkedList"

	"go.zhuzi.me/go/log"
)

func main() {
	head := new(linkedList.ListNode)
	head.Val = 1
	ln2 := new(linkedList.ListNode)
	ln2.Val = 2
	ln3 := new(linkedList.ListNode)
	ln3.Val = 3
	ln4 := new(linkedList.ListNode)
	ln4.Val = 4
	ln5 := new(linkedList.ListNode)
	ln5.Val = 5

	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	p := linkedList.ReverseList(head)
	log.Error(p)
}
