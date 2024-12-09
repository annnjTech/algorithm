package comp

import (
	"math"
	"math/rand"
	"time"
)

// LenRandomValueRandom 返回一个切片arr，arr长度[0, maxLen - 1]，arr中的每个值[1, maxValue - 1]
func LenRandomValueRandom(maxLen, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := r.Intn(maxLen) + 1
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = r.Intn(maxValue) + 1
	}
	return arr
}

// CopyArray 复制一个切片arr
func CopyArray(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

// IsSorted arr1和arr2一定等长
func IsSorted(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	maxValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxValue > arr[i] {
			return false
		}
		maxValue = (int)(math.Max((float64)(maxValue), (float64)(arr[i])))
	}
	return true
}
