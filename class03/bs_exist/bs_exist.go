package bs_exist

// Find 二分查找有序数组arr中是否存在num
func Find(arr []int, num int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == num {
			return true
		} else if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
