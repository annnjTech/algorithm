package main

import (
	"algorithm/class06/priority_queue"
	"container/heap"
)

func main() {
	// 创建一个优先队列
	pq := make(priority_queue.PriorityQueue, 0)
	heap.Init(&pq)

	// 插入元素
	heap.Push(&pq, &priority_queue.Student{Name: "张三", Age: 26, Id: 5})
	heap.Push(&pq, &priority_queue.Student{Name: "李四", Age: 18, Id: 1})
	heap.Push(&pq, &priority_queue.Student{Name: "王五", Age: 22, Id: 4})
	heap.Push(&pq, &priority_queue.Student{Name: "赵六", Age: 15, Id: 2})
	heap.Push(&pq, &priority_queue.Student{Name: "许七", Age: 34, Id: 3})

	for pq.Len() > 0 {
		s := heap.Pop(&pq).(*priority_queue.Student)
		println(s.Name, s.Age, s.Id)
	}
}
