package priority_queue

type Student struct {
	Name string
	Age  int
	Id   int
}

type PriorityQueue []*Student

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	// 我们希望年龄大的学生排在前面，所以使用大于号
	return p[i].Id > p[j].Id
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	student := x.(*Student)
	*p = append(*p, student)
}

func (p *PriorityQueue) Pop() interface{} {
	length := len(*p)
	student := (*p)[length-1]
	(*p)[length-1] = nil // 避免内存泄漏
	*p = (*p)[0 : length-1]
	return student
}
