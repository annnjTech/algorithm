package merge_k_sorted_list

import (
	"algorithm/class06/model"
	"algorithm/class06/priority_queue"
	"container/heap"
)

func MergeKLists(lists []*model.Node) *model.Node {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	pq := make(priority_queue.PriorityQueue, 0)
	heap.Init(&pq)

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&pq, lists[i])
		}
	}

	if pq.Len() == 0 {
		return nil
	}

	var head, pre *model.Node
	head = heap.Pop(&pq).(*model.Node)
	pre = head
	if pre.Next != nil {
		heap.Push(&pq, pre.Next)
	}
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*model.Node)
		pre.Next = cur
		pre = cur
		if cur.Next != nil {
			heap.Push(&pq, cur.Next)
		}
	}
	return head
}
