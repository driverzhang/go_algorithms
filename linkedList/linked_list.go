package linkedList

// 单链表的基本操作

import "fmt"

/*
单链表基本操作
author:leo
*/

type SListNode struct {
	next  *SListNode
	value interface{}
}

type LinkedList struct {
	head   *SListNode // 表示头指针
	length uint
}

func NewSListNode(v interface{}) *SListNode {
	return &SListNode{nil, v}
}

func (this *SListNode) GetNext() *SListNode {
	return this.next
}

func (this *SListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *SListNode {
	return &SListNode{NewSListNode(0), 0}
}

//在p节点 后面 插入v值的节点
func (this *LinkedList) InsertAfter(p *SListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewSListNode(v) // 新建一个节点并赋值
	oldNext := p.next          // 取出 p 原本指向的下一个节点
	p.next = newNode           // 将节点P 指向 新建节点
	newNode.next = oldNext     // 把 新建节点 指向 之前p原本指向的下一个节点
	this.length++
	return true
}

//在p节点 前面 插入v值的节点
func (this *LinkedList) InsertBefore(p *SListNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.next // 从头指针还是找 直到找到 cur == p 点为止
	pre := this.head

	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}

	newNode := NewSListNode(v) // 此时的 cur == p ,pre 是 cur 前一节点
	pre.next = newNode         // 要将 新节点插入到 pre 和 cur 中间即可
	newNode.next = cur
	this.length++
	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next // 一直跑下去
	}
	return this.InsertAfter(cur, v)
}

//通过 索引 查找节点
func (this *LinkedList) FindByIndex(index uint) *SListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除 传入的节点
func (this *LinkedList) DeleteNode(p *SListNode) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next // 改变相邻前后指针关系
	p = nil           // 释放内存
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
