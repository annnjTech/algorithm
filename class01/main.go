package main

import (
	"algorithm/class01/algorithm"
	"fmt"
)

func main() {
	arr := []int{7, 1, 3, 4, 2, 8, 3, 1, 9, 4, 5}
	fmt.Println(arr)
	algorithm.InsertSort2(arr)
	fmt.Println(arr)
}
