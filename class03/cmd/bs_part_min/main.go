package main

import "algorithm/class03/bs_part_min"

func main() {
	println("测试开始")
	arr := []int{4, 2, 3, 1, 5}
	ans := bs_part_min.BsPartMinIndex(arr)
	println("ans:", ans)

	/*maxLen := 100
	maxValue := 200
	testTimes := 1000000
	for i := 0; i < testTimes; i++ {
		arr := comp.RandomArray(maxLen, maxValue)
		ans := bs_part_min.BsPartMinIndex(arr)
		if !comp.TestBsPartMin(arr, ans) {
			common.PrintArray(arr)
			println("ans:", ans)
			break
		}
	}*/
	println("测试结束")
}
