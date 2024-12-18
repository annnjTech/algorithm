package bs_part_min

// BsPartMinIndex 返回局部最小值的索引
func BsPartMinIndex(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	N := len(arr)
	if N == 1 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[N-1] < arr[N-2] {
		return N - 1
	}
	left, right := 0, N-1
	// left ... right 肯定有局部最小
	for left < right-1 {
		mid := (left + right) / 2
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
			return mid
		} else {
			if arr[mid] > arr[mid-1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if arr[left] < arr[right] {
		return left
	} else {
		return right
	}
}
