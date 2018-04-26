package main

import (

)

type Node struct {
	Value 		interface{}
	next, prev  *Node
}

type List struct {
	root 	Node
	length 	int
}

func New1() *List  {
	l := &List{}
	l.length = 0
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

func (l *List) IsEmpty() bool {
	return l.root.next == &l.root
}

func (l *List) Length() int  {
	return l.length
}

func (l *List) PushFront(elements ...interface{})  {
	for _, element := range elements {
		n := &Node{Value:element}
		n.next = l.root.next
		n.prev = &l.root
		l.root.next.prev = n
		l.root.next = n
		l.length++
	}
}

func (l *List) PushBack(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = &l.root     // since n is the last element, its next should be the head
		n.prev = l.root.prev // n's prev should be the tail
		l.root.prev.next = n // tail's next should be n
		l.root.prev = n      // head's prev should be n
		l.length++
	}
}

func (l *List) Find(element interface{}) int {
	index := 0
	p := l.root.next
	for p != &l.root && p.Value !=element {
		p = p.next
		index++
	}

	if p != &l.root {
		return index
	}
	return -1
}

func (l *List) Remove(n *Node)  {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	l.length--
}

func (l *List) Lpop() interface{} {
	if l.length == 0 {
		return nil
	}
	n := l.root.next
	l.Remove(n)
	return n.Value
}

func (l *List) normalIndex(index int) int  {
	if index > l.length-1 {
		index = l.length - 1
	}
	if index < -l.length {
		index = 0
	}

	index = (l.length + index)%l.length
	return index
}

func (l *List)index(a int) *Node  {
	return l.index(a)
}

func (l *List) Range(start, end int) []interface{} {
	// 获取正常的start和end
	start = l.normalIndex(start)
	end = l.normalIndex(end)
	// 声明一个interface类型的数组
	res := []interface{}{}
	// 如果上下界不符合逻辑，返回空res
	if start > end {
		return res
	}

	sNode := l.index(start)
	eNode := l.index(end)
	// 起始点和重点遍历
	for n := sNode; n != eNode; {
		// res的append方式
		res = append(res, n.Value)
		n = n.next
	}
	res = append(res, eNode.Value)
	return res
}
