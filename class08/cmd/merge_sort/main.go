package main

import (
	"algorithm/class08/merge_sort"
	"fmt"
)

func main() {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	merge_sort.MergeSort2(&arr)
	fmt.Println(arr)
}
