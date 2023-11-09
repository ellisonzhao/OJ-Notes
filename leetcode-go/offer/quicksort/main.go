package main

import (
	"fmt"
)

func MySort(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return arr
	}
	// write code heres
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, left, right int) {
	if left < right {
		pos := partition(arr, left, right)
		quickSort(arr, left, pos-1)
		quickSort(arr, pos+1, right)
	}
}

func partition(list []int, low, high int) int {
	pivot := list[low]
	// 导致 low 位置值为空
	for low < high {
		// high 指针值 >= pivot high 指针左移
		for low < high && pivot <= list[high] {
			high--
		}
		// 填补 low 位置空值
		// high 指针值 < pivot high 值 移到 low 位置
		// high 位置值空
		list[low] = list[high]
		// low 指针值 <= pivot low 指针右移
		for low < high && pivot >= list[low] {
			low++
		}
		// 填补 high 位置空值
		// low 指针值 > pivot low 值 移到 high 位置
		// low 位置值空
		list[high] = list[low]
	}
	// pivot 填补 low 位置的空值
	list[low] = pivot
	return low
}

func main() {
	arr := []int{5, 2, 3, 1, 4}
	fmt.Println(MySort(arr))
}
