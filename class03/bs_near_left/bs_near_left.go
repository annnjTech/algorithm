package bs_near_left

// MostLeftNoLessNumIndex 返回有序数组 arr 中第一个不小于 num 的元素的索引
func MostLeftNoLessNumIndex(arr []int, num int) int {
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
