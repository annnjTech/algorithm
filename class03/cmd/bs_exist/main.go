package main

import (
	"algorithm/class03/bs_exist"
	"algorithm/class03/bs_exist/comp"
	"fmt"
	"math/rand"
	sort "sort2"
	"time"
)

func main() {
	maxSize := 10
	maxValue := 100
	testTimes := 500000
	success := true
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < testTimes; i++ {
		arr := comp.GenerateRandomArray(maxSize, maxValue)
		sort.InsertSort(arr)
		value := (int)((float64)(maxValue+1)*r.Float64()) - (int)((float64)(maxValue)*r.Float64())
		if comp.Test(arr, value) != bs_exist.Find(arr, value) {
			fmt.Println("出错啦！")
			success = false
			break
		}
	}
	if success {
		fmt.Println("测试通过！")
	} else {
		fmt.Println("测试失败！")
	}
}
