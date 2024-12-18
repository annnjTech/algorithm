package main

import (
	"algorithm/class01/sort"
	"algorithm/class03/bs_near_right"
	"algorithm/class03/bs_near_right/comp"
	"algorithm/class03/common"
	"math/rand"
	"time"
)

func main() {
	maxSize := 10
	maxValue := 100
	testTimes := 500000
	success := true
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < testTimes; i++ {
		arr := common.GenerateRandomArray(maxSize, maxValue)
		sort.InsertSort2(arr)
		value := (int)((float64)(maxValue+1)*r.Float64()) - (int)((float64)(maxValue)*r.Float64())
		if comp.TestBsNearRight(arr, value) != bs_near_right.SmallerThanNumMostRightIndex(arr, value) {
			common.PrintArray(arr)
			println("value:", value)
			println(comp.TestBsNearRight(arr, value))
			println(bs_near_right.SmallerThanNumMostRightIndex(arr, value))
			success = false
			break
		}
	}
	if success {
		println("All tests pass!")
	} else {
		println("Some tests fail!")
	}
}
