package bs_near_left

// BiggerThanNumMostLeftIndex 返回有序数组arr中>=num的最左索引，如果不存在，则返回-1
func BiggerThanNumMostLeftIndex(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	left, right := 0, len(arr)-1
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] >= num {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
