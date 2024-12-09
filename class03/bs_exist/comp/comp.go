package comp

import (
	"math/rand"
	"time"
)

func Test(sortedArr []int, num int) bool {
	for _, curr := range sortedArr {
		if curr == num {
			return true
		}
	}
	return false
}

func GenerateRandomArray(maxSize, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, (int)((float64)(maxSize+1)*r.Float64()))
	for i := 0; i < len(arr); i++ {
		arr[i] = (int)((float64)(maxValue+1)*r.Float64()) - (int)((float64)(maxValue)*r.Float64())
	}
	return arr
}
