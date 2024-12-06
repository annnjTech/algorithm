package main

import (
	"algorithm/class01/sort"
)

func main() {
	arr := []int{7, 1, 3, 4, 2, 8, 3, 1, 9, 4, 5}
	sort.PrintArray(arr)
	sort.InsertSort2(arr)
	sort.PrintArray(arr)
}
