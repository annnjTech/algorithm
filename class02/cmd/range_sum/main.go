package main

import (
	"algorithm/class02/rang_sum"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(rang_sum.RangeSum1(arr, 2, 4))
	fmt.Println(rang_sum.RangeSum2(arr, 2, 4))
}
