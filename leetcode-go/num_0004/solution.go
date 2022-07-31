package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 4, 9}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	left := (l1 + l2 + 1) / 2
	right := (l1+l2)/2 + 1
	// 将偶数和奇数的情况合并，如果是奇数，会求两次同样的 k
	first := getKth(nums1, 0, l1, nums2, 0, l2, left)
	fmt.Println("first: ", first)
	second := getKth(nums1, 0, l1, nums2, 0, l2, right)
	fmt.Println("second: ", second)

	return float64(first+second) / 2
}

// 找到第 k 小的数
func getKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
	l1 := end1 - start1
	l2 := end2 - start2
	// 这样能保证如果有数组为空，一定是 nums1 为空
	if l1 > l2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	// 边界条件
	if l1 == 0 {
		// 相当于 start2 位置是第 1 小的数，那第 k 小的数就在 start2+k-1 处
		return nums2[start2+k-1]
	}

	// 递归终止条件
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	// nums1 中第 k/2 小的数所在位置
	i := start1 + min(l1, k/2) - 1
	// nums2 中第 k/2 小的数所在位置
	j := start2 + min(l2, k/2) - 1
	if nums1[i] < nums2[j] {
		// 排除 nums1 中 [start1, i] 的元素
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}

	// 排除 nums2 中 [start2, j] 的元素
	return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
