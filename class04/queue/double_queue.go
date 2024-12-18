package queue

import "algorithm/class04/model"

type MyDoubleQueue struct {
	Head *model.DoubleNode
	Tail *model.DoubleNode
	Len  int
}

// IsEmpty 判断队列是否为空
func (q *MyDoubleQueue) IsEmpty() bool {
	return q.Len == 0
}

// Length 返回队列长度
func (q *MyDoubleQueue) Length() int {
	return q.Len
}

// PushHead 从队列头部添加元素
func (q *MyDoubleQueue) PushHead(value int) {
	cur := new(model.DoubleNode)
	cur.Value = value
	if q.Head == nil {
		q.Head = cur
		q.Tail = cur
	} else {
		cur.Next = q.Head
		q.Head.Last = cur
		q.Head = cur
	}
	q.Len++
}

// PushTail 从队列尾部添加元素
func (q *MyDoubleQueue) PushTail(value int) {
	cur := new(model.DoubleNode)
	cur.Value = value
	if q.Head == nil {
		q.Head = cur
		q.Tail = cur
	} else {
		q.Tail.Next = cur
		cur.Last = q.Tail
		q.Tail = cur
	}
	q.Len++
}

// PollHead 从队列头部删除元素
func (q *MyDoubleQueue) PollHead() int {
	val := 0
	if q.Head == nil {
		return val
	}
	q.Len--
	val = q.Head.Value
	if q.Head == q.Tail {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Next
		q.Head.Last = nil
	}
	return val
}

// PollTail 从队列尾部删除元素
func (q *MyDoubleQueue) PollTail() int {
	val := 0
	if q.Head == nil {
		return val
	}
	q.Len--
	val = q.Tail.Value
	if q.Head == q.Tail {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Tail = q.Tail.Last
		q.Tail.Next = nil
	}
	return val
}

// PeekHead 获取但不删除队列头部元素
func (q *MyDoubleQueue) PeekHead() int {
	val := 0
	if q.Head != nil {
		val = q.Head.Value
	}
	return val
}

// PeekTail 获取但不删除队列尾部元素
func (q *MyDoubleQueue) PeekTail() int {
	val := 0
	if q.Tail != nil {
		val = q.Tail.Value
	}
	return val
}
