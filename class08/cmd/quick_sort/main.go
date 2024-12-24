package main

import (
	"algorithm/class08/quick_sort"
	"fmt"
)

func main() {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	//quick_sort.QuickSort1(arr)
	quick_sort.QuickSort2(arr)
	fmt.Println(arr)
}
