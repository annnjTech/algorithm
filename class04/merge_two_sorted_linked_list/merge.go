package merge_two_sorted_linked_list

import "algorithm/class04/model"

// MergeTwoSortedLinkedLists 合并两个有序链表，返回合并后的链表的头结点
func MergeTwoSortedLinkedLists(head1, head2 *model.Node) *model.Node {
	if head1 == nil || head2 == nil {
		if head1 == nil {
			return head2
		} else {
			return head1
		}
	}
	var head, cur1, cur2, pre *model.Node
	// 确定两个链表中的头结点中的最小值，作为合并后的链表的头结点
	if head1.Value <= head2.Value {
		head = head1
	}

	// cur1用于遍历链表1
	// cur2用于遍历链表2
	cur1 = head.Next
	if head == head1 {
		cur2 = head2
	} else {
		cur2 = head1
	}
	// pre用于指向合并后的链表的尾结点
	pre = head
	for cur1 != nil && cur2 != nil {
		if cur1.Value <= cur2.Value {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}
	// 处理链表1和链表2剩余的结点
	if cur1 != nil {
		pre.Next = cur1
	} else {
		pre.Next = cur2
	}
	return head
}

/*
head1: ①->③->⑤->⑦->⑨->nil
head2: ②->④->⑥->nil
合并后的链表: ①->②->③->④->⑤->⑥->⑦->⑨->nil
*/
