package main

import (
	"algorithm/class01/sort"
	"algorithm/class03/bs_exists"
	"algorithm/class03/bs_exists/comp"
	"algorithm/class03/common"
	"fmt"
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
		if comp.TestBsExists(arr, value) != bs_exists.BsFind(arr, value) {
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
