package merge_k_sorted_list

import "algorithm/class06/model"

type PriorityQueue []*model.Node

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	// 我们希望值小的的元素排在前面，所以使用小于号
	return p[i].Value < p[j].Value
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	student := x.(*model.Node)
	*p = append(*p, student)
}

func (p *PriorityQueue) Pop() interface{} {
	length := len(*p)
	node := (*p)[length-1]
	(*p)[length-1] = nil // 避免内存泄漏
	*p = (*p)[0 : length-1]
	return node
}
