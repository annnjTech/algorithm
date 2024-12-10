package bs_near_right

// SmallerThanNumMostRightIndex 返回有序数组arr中<=num的最右索引，如果不存在，则返回-1
func SmallerThanNumMostRightIndex(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	left, right := 0, len(arr)-1
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] <= num {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
