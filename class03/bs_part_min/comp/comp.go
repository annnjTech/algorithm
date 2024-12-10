package comp

import (
	"math/rand"
	"time"
)

func TestBsPartMin(arr []int, minIndex int) bool {
	if arr == nil {
		return false
	}
	if len(arr) == 0 {
		return minIndex == -1
	}
	left := minIndex - 1
	right := minIndex + 1
	var leftBigger, rightBigger bool
	if left >= 0 {
		leftBigger = arr[left] > arr[minIndex]
	} else {
		leftBigger = true
	}
	if right < len(arr) {
		rightBigger = arr[right] > arr[minIndex]
	} else {
		rightBigger = true
	}

	return leftBigger && rightBigger
}

func RandomArray(maxLen, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, (int)((float64)(maxLen+1)*r.Float64()))
	if len(arr) > 0 {
		arr[0] = r.Intn(maxValue)
		for i := 1; i < len(arr); i++ {
			arr[i] = r.Intn(maxValue)
			for arr[i] == arr[i-1] {
				arr[i] = r.Intn(maxValue)
			}
		}
	}
	return arr
}
