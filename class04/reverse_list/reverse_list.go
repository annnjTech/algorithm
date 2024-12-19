package reverse_list

import "algorithm/class04/model"

// ReverseLinkedList 反转单链表
/**
初始状态：
       ① -> ② -> ③ -> nil
反转后：
nil <- ③ <- ② <- ①
1) tmpPrev = nil, tmpNext = nil
2) 判断head是否为空，为空则退出循环
3) head指向①节点，tmpNext指向下一个节点，即②节点
4) head的next指向tmpPrev，即head的next指向nil
5) tmpPrev指向head
6) head指向tmpNext指向的节点，即head指向②节点
7) 重复步骤2-6，直到head指向nil
8) 返回tmpPrev，此时的tmpPrev指向了③，即返回①节点
*/
func ReverseLinkedList(head *model.Node) *model.Node {
	var tmpPrev, tmpNext *model.Node
	for head != nil {
		tmpNext = head.Next
		head.Next = tmpPrev
		tmpPrev = head
		head = tmpNext
	}
	return tmpPrev
}

// ReverseDoubleLinkedList 反转双向链表
/**
初始状态：
	    ① -> ② -> ③ -> nil                 next线条指向
nil  <- ① <- ② <- ③                        last线条指向
反转后：
        ③ -> ② -> ① -> nil                 next线条指向
nil  <- ③ <- ② <- ①                        last线条指向
1) tmpPrev = nil, tmpNext = nil
2) 判断head是否为空，为空则退出循环
3) head指向①节点，tmpNext指向head的next，即②节点
4) head的next指向tmpPrev，即head的next指向nil
5) head的last指向tmpNext，即head的last指向②节点
6) tmpPrev指向head，即tmpPrev指向①节点
7) head指向tmpNext，即head指向②节点
8) 重复步骤2-7，直到head指向nil
9) 返回tmpPrev，此时的tmpPrev指向了③，即返回③节点
*/
func ReverseDoubleLinkedList(head *model.DoubleNode) *model.DoubleNode {
	var tmpPrev, tmpNext *model.DoubleNode
	for head != nil {
		tmpNext = head.Next
		head.Next = tmpPrev
		head.Last = tmpNext
		tmpPrev = head
		head = tmpNext
	}
	return tmpPrev
}
