package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 0}
	fmt.Println(reversePairs(nums))
}

func reversePairs(nums []int) int {
	return merge(nums, 0, len(nums)-1)
}

// 归并排序思想
func merge(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2
	// 先递归处理好左半部分和右半部分
	res := merge(nums, left, mid) + merge(nums, mid+1, right)
	// 计算当前的逆序对
	i, j := left, mid+1
	var tmpNums []int
	for i <= mid && j <= right {
		// 形成逆序对
		if nums[i] > nums[j] {
			res += mid - i + 1
			tmpNums = append(tmpNums, nums[j])
			j++
		} else {
			tmpNums = append(tmpNums, nums[i])
			i++
		}
	}
	for i <= mid {
		tmpNums = append(tmpNums, nums[i])
		i++
	}
	for j <= right {
		tmpNums = append(tmpNums, nums[j])
		j++
	}

	k := left
	for _, t := range tmpNums {
		nums[k] = t
		k++
	}

	return res
}
