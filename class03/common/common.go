package common

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(maxSize, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, (int)((float64)(maxSize+1)*r.Float64()))
	for i := 0; i < len(arr); i++ {
		arr[i] = (int)((float64)(maxValue+1)*r.Float64()) - (int)((float64)(maxSize)*r.Float64())
	}
	return arr
}

func PrintArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()
}
