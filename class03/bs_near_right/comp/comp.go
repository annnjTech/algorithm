package comp

func TestBsNearRight(arr []int, num int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= num {
			return i
		}
	}
	return -1
}
