package sort

import "fmt"

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// SelectSort 选择排序
func SelectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	length := len(arr)

	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			swap(arr, i, minIndex)
		}
	}
}

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	length := len(arr)

	for end := length - 1; end >= 0; end-- {
		for second := 1; second <= end; second++ {
			if arr[second-1] > arr[second] {
				swap(arr, second-1, second)
			}
		}
	}
}

// InsertSort1 插入排序1
func InsertSort1(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	length := len(arr)

	for i := 1; i < length; i++ {
		newNumIndex := i
		for ; newNumIndex-1 >= 0 && arr[newNumIndex-1] > arr[newNumIndex]; newNumIndex-- {
			swap(arr, newNumIndex, newNumIndex-1)
		}
	}
}

// InsertSort2 插入排序2
func InsertSort2(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	length := len(arr)

	for i := 1; i < length; i++ {
		for pre := i - 1; pre >= 0 && arr[pre] > arr[pre+1]; pre-- {
			swap(arr, pre, pre+1)
		}
	}
}

func PrintArray(arr []int) {
	for _, item := range arr {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
