package queue

import "algorithm/class04/model"

type MyQueue struct {
	Head *model.Node
	Tail *model.Node
	Len  int
}

// IsEmpty 队列是否为空
func (q *MyQueue) IsEmpty() bool {
	return q.Len == 0
}

// Length 队列长度
func (q *MyQueue) Length() int {
	return q.Len
}

// Offer 添加元素到队列尾部
func (q *MyQueue) Offer(val int) {
	cur := new(model.Node)
	cur.Value = val
	if q.Tail == nil {
		q.Head = cur
		q.Tail = cur
	} else {
		q.Tail.Next = cur
		q.Tail = cur
	}
	q.Len++
}

// Poll 移除队列头部元素并返回
func (q *MyQueue) Poll() int {
	val := 0
	if q.Head != nil {
		val = q.Head.Value
		q.Head = q.Head.Next
	}
	if q.Head == nil {
		q.Tail = nil
	}
	q.Len--
	return val
}

// Peek 获取但不移除队列头部元素
func (q *MyQueue) Peek() int {
	if q.Head != nil {
		return q.Head.Value
	}
	return 0
}
