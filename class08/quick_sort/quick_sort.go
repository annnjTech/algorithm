package quick_sort

// QuickSort1 快速排序1，递归的方式实现，时间复杂度：O(nlogn)
func QuickSort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	process(&arr, 0, len(arr)-1)
}

func process(arr *[]int, left, right int) {
	if left >= right {
		return
	}
	// equalFirstIndex，表示=区域的第一个元素的下标
	// equalLastIndex，表示=区域的最后一个元素的下标
	equalFirstIndex, equalLastIndex := partition(arr, left, right)
	process(arr, left, equalFirstIndex-1) // 递归处理 < 区域，从left...equalFirstIndex-1
	process(arr, equalLastIndex+1, right) // 递归处理 > 区域，从equalLastIndex+1...right
}

// partition，分区函数，把数组arr[left...right]分成<P的区域，=P的区域，>P的区域
/*
一组L...R的元素中，以最后一个元素P值为标准，分成三个区域：<P的区域在左边，=P的区域在中间，>P的区域在右边
P表示数组的最后一个元素的值，下标是right
lessR表示<区域的右边届，在最开始0~len(arr)-1时，lessR指向数组的-1位置，实质上此时<区域为空
moreL表示>区域的左边界，在最开始0~len(arr)-1时，moreL指向数组的最后一个元素

1.当前index指向的元素的值<P的值，当前index和lessR指向的元素交换位置，lessR往右扩1位，index往右移1位
2.当前index指向的元素的值>P的值，当前index和moreL指向的元素交换位置，moreL往左扩1位，index不动
3.当前index指向的元素的值=P的值，index往右移1位
退出循环的条件：当index移动到moreL的最左边的时候
最后还要把P的位置right和moreL指向的元素交换位置，即把最后一个元素的位置和>区域的左边界的元素交换位置
最终我们返回=区域的第一个元素的下标和=区域的最后一个元素的下标
*/
func partition(arr *[]int, left, right int) (equalFirstIndex, equalLastIndex int) {
	lessR := left - 1 // 记录 < 区域的右边届
	moreL := right    // 记录 > 区域的左边界，把最后一个元素先纳入到 > 区域
	index := left     // index指向当前处理的元素
	for index < moreL {
		if (*arr)[index] < (*arr)[right] {
			swap(arr, lessR+1, index)
			lessR++
			index++
		} else if (*arr)[index] > (*arr)[right] {
			swap(arr, moreL-1, index)
			moreL--
		} else {
			index++
		}
	}
	swap(arr, moreL, right)
	return lessR + 1, moreL
}

func swap(arr *[]int, i, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

// QuickSort2 快速排序2，非递归的方式实现，时间复杂度：O(nlogn)
func QuickSort2(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 栈的元素是 边界数组[left, right]，表示待排序的下标范围
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{0, len(arr) - 1}) // 边界数组最开始是[0, len(arr)-1]，入栈处理
	for len(stack) > 0 {                           // 栈不为空时
		top := stack[len(stack)-1]   // 获取栈顶元素
		stack = stack[:len(stack)-1] // 栈顶元素出栈
		// equalFirstIndex，表示=区域的第一个元素的下标
		// equalLastIndex，表示=区域的最后一个元素的下标
		equalFirstIndex, equalLastIndex := partition(&arr, top[0], top[1])
		if top[0] < equalFirstIndex { // 有 < 区域，left...equalFirstIndex-1的边界数组入栈
			stack = append(stack, [2]int{top[0], equalFirstIndex - 1})
		}
		if equalLastIndex < top[1] { // 有 > 区域，equalLastIndex+1...right的边界数组入栈
			stack = append(stack, [2]int{equalLastIndex + 1, top[1]})
		}
	}
}
