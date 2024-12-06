package rang_sum

func RangeSum1(arr []int, left, right int) (sum int) {
	for i := left; i <= right; i++ {
		sum += arr[i]
	}
	return
}

func RangeSum2(arr []int, left, right int) (sum int) {
	n := len(arr)
	// 前缀和 数组
	pre := make([]int, n)
	pre[0] = arr[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] + arr[i]
	}

	if left == 0 {
		sum = pre[right]
	} else {
		sum = pre[right] - pre[left-1]
	}
	return
}
