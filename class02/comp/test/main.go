package main

import (
	"algorithm/class01/sort"
	"algorithm/class02/comp"
	"fmt"
)

func main() {
	maxLen := 5
	maxValue := 1000
	testTimes := 10000
	for i := 0; i < testTimes; i++ {
		arr := comp.LenRandomValueRandom(maxLen, maxValue)
		tmp := comp.CopyArray(arr)
		sort.BubbleSort(arr)
		if !comp.IsSorted(arr) {
			for j := 0; j < len(tmp); j++ {
				fmt.Print(tmp[j], " ")
			}
			fmt.Println()
			fmt.Println("选择排序错了！")
			break
		}
	}
}
