package merge_sort

// MergeSort1 递归实现归并排序，时间复杂度：O(nlogn)
func MergeSort1(arr *[]int) {
	length := len(*arr)
	if arr == nil || length < 2 {
		return
	}
	process(arr, 0, length-1)
}

func process(arr *[]int, l, r int) {
	if l == r {
		return
	}
	// mid := (l + r) / 2，为什么不用这种方式计算mid呢？因为这样当数值很大的时候，相加的时候有可能会溢出
	mid := l + ((r - l) >> 1)
	process(arr, l, mid)
	process(arr, mid+1, r)
	merge(arr, l, mid, r)
}

// MergeSort2 非递归实现归并排序，时间复杂度：O(nlogn)
/* 原数组：1 0 9 8 3 2 11 10 5 4 13 12 7 6 14
========================================
step = 1
合并前分组：
l1: [1]   r1: [0]
l2: [9]   r2: [8]
l3: [3]   r3: [2]
l4: [11]  r4: [10]
l5: [5]   r5: [4]
l6: [13]  r6: [12]
l7: [7]   r7: [6]
l8: [14]
合并后分组：
[0 1] [8 9] [2 3] [10 11] [4 5] [12 13] [6 7] [14]
mid == length-1，break此次for l < length的循环
step > 15/2吗？否，所以继续下一轮for step < length的循环
========================================
step = 2
合并前分组：
l1: [0 1]     r1: [8 9]
l2: [2 3]     r2: [10 11]
l3: [4 5]     r3: [12 13]
l4: [6 7]     r4: [14]
合并后分组：
[0 1 2 3] [8 9 10 11] [4 5 6 7] [12 13 14]
r == length-1，break此次for l < length的循环
step > 15/2吗？否，所以继续下一轮for step < length的循环
========================================
step = 4
合并前分组：
l1: [0 1 2 3]   r1: [8 9 10 11]
l2: [4 5 6 7]   r2: [12 13 14]
合并后分组：
[0 1 2 3 4 5 6 7] [8 9 10 11 12 13 14]
r == length-1，break此次for l < length的循环
step > 15/2吗？否，所以继续下一轮for step < length的循环
========================================
step = 8
合并前分组：
l1: [0 1 2 3 4 5 6 7]   r1: [8 9 10 11 12 13 14]
合并后分组：
[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
r == length-1，break此次for l < length的循环
step > 15/2吗？是，结束for step < length的循环
合并排序结束*/
func MergeSort2(arr *[]int) {
	length := len(*arr)
	if arr == nil || length < 2 {
		return
	}

	step := 1
	for step < length {
		l := 0
		for l < length {
			var mid int
			// 左边组
			// 如果下标length-1到l的距离(即length-1-l+1)还够步长step，则mid=从l开始包括l自己数step个元素的最大下标，即l+step-1
			// 否则mid就数到下标length-1，即数到最后一个元素
			if length-1-l+1 >= step {
				mid = l + step - 1
			} else {
				mid = length - 1
			}
			// 如果mid已经到达最后一个元素，则不需要再进行本轮的merge操作了
			if mid == length-1 {
				break
			}
			var r int
			// 右边组
			// 如果下标length-1到mid+1的距离(length-1-(mid+1)+1)还够步长step，则右边届r=从mid+1开始包括mid+1自己数step个元素的最大下标，即(mid+1)+step-1
			// 否则r就数到下标length-1，即数到最后一个元素
			if length-1-(mid+1)+1 >= step {
				r = (mid + 1) + step - 1
			} else {
				r = length - 1
			}
			// 能来到这里说明左边组和右边组都有元素，进行merge操作
			merge(arr, l, mid, r)
			// 如果r已经到达最后一个元素了，则后面没有左右两边组了，可以退出本轮的循环
			if r == length-1 {
				break
			}
			// l的位置移动到r+1的位置，准备下一个左右组的merge操作
			l = r + 1
		}
		if step > length/2 { // 这里不能>=，因为假设当step=8，总长度是17时，step=8 >= 17/2=8，跳出循环了，就无法进行step=16的时候，进行左边16个数和右边1个数的merge操作
			break
		}
		step <<= 1 // 步长*2，步长为2的幂
	}
}

func merge(arr *[]int, l, mid, r int) {
	help := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := mid + 1
	for p1 <= mid && p2 <= r {
		if (*arr)[p1] <= (*arr)[p2] {
			help[i] = (*arr)[p1]
			p1++
		} else {
			help[i] = (*arr)[p2]
			p2++
		}
		i++
	}
	// 要么p1越界，要么p2越界
	// 不可能出现共同越界
	for p1 <= mid {
		help[i] = (*arr)[p1]
		p1++
		i++
	}
	for p2 <= r {
		help[i] = (*arr)[p2]
		p2++
		i++
	}
	for i = 0; i < len(help); i++ {
		(*arr)[l+i] = help[i]
	}
}
