package add_two_linked_list

import "algorithm/class04/model"

func AddTwoLinkedList(head1, head2 *model.Node) *model.Node {
	len1, len2 := ListLength(head1), ListLength(head2)
	// 设定长链表为long，短链表为short
	var long, short *model.Node
	if len1 >= len2 {
		long = head1
	} else {
		long = head2
	}
	if long == head1 {
		short = head2
	} else {
		short = head1
	}

	// curLong，用于遍历long链表，curLong我们作为最后的结果链表，省去了新建链表的操作
	// curShort，用于遍历short链表
	var curLong, curShort *model.Node
	curLong = long
	curShort = short

	var carry int     // 用于记录进位信息
	var curNumber int // 用于记录当前位相加结果

	var lastNode *model.Node // 用于记录最后一个节点，用于如果最后一位有进位，需要新建节点
	for curShort != nil {
		curNumber = curLong.Value + curShort.Value + carry
		curLong.Value = curNumber % 10 // 当前节点的值等于当前位相加结果的余数
		carry = curNumber / 10         // 进位信息等于当前位相加结果的商
		lastNode = curLong
		curLong = curLong.Next
		curShort = curShort.Next
	}
	for curLong != nil {
		curNumber = curLong.Value + carry
		curLong.Value = curNumber % 10 // 当前节点的值等于当前位相加结果的余数
		carry = curNumber / 10         // 进位信息等于当前位相加结果的商
		lastNode = curLong
		curLong = curLong.Next
	}
	// 最后一位有进位，需要新建节点
	if carry != 0 {
		lastNode.Next = &model.Node{Value: 1}
	}
	return long
}

func ListLength(head *model.Node) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

/**
head1: ①->②->③->④->⑨->nil
head2: ①->⑤->⑨->⑥->nil
结果：②->⑦->②->①->◎->①->nil，最后一位有进位
*/
