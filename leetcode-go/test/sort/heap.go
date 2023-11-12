package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 6, 3, 7, 2, 8, 1, 4, 5}
	fmt.Println("Before sorting:", arr)
	heapSort(arr)
	fmt.Println("After sorting:", arr)
}

func heapSort(arr []int) {
	n := len(arr)
	// 构建大顶堆
	for i := n/2 - 1; i >= 0; i-- {
		adjustHeap(arr, n, i)
	}

	// 逐个提取堆顶元素，与最后一个元素交换，并调整堆结构
	for i := n - 1; i > 0; i-- {
		// 交换堆顶元素和最后一个元素
		arr[0], arr[i] = arr[i], arr[0]

		// 调整堆结构
		adjustHeap(arr, i, 0)
	}
}

func adjustHeap(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// 如果左子节点大于根节点
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点大于当前最大节点
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大节点不是根节点
	if largest != i {
		// 交换当前节点和最大节点
		arr[i], arr[largest] = arr[largest], arr[i]

		// 递归调整子堆
		adjustHeap(arr, n, largest)
	}
}
