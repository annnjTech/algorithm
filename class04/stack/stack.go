package stack

import "algorithm/class04/model"

type MyStack struct {
	Head *model.Node
	Len  int
}

// IsEmpty 栈是否为空
func (s *MyStack) IsEmpty() bool {
	return s.Len == 0
}

// Length 栈的长度
func (s *MyStack) Length() int {
	return s.Len
}

// Push 压入元素
func (s *MyStack) Push(value int) {
	cur := new(model.Node)
	cur.Value = value
	if s.Head == nil {
		s.Head = cur
	} else {
		cur.Next = s.Head
		s.Head = cur
	}
	s.Len++
}

// Pop 弹出栈顶元素
func (s *MyStack) Pop() int {
	val := 0
	if s.Head != nil {
		val = s.Head.Value
		s.Head = s.Head.Next
		s.Len--
	}
	return val
}

// Peek 获取但不弹出栈顶元素
func (s *MyStack) Peek() int {
	if s.Head != nil {
		return s.Head.Value
	}
	return 0
}
