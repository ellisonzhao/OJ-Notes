package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow2(nums, k))
}

func maxSlidingWindow2(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || k <= 0 {
		return nil
	}

	n := len(nums)
	maxLeft := make([]int, n)
	maxLeft[0] = nums[0]
	maxRight := make([]int, n)
	maxRight[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		if i%k == 0 {
			// 新分块中第一个元素
			maxLeft[i] = nums[i]
		} else {
			maxLeft[i] = max(maxLeft[i-1], nums[i])
		}

		j := n - i - 1
		if j%k == k-1 {
			// 新分块中最后一个元素
			maxRight[j] = nums[j]
		} else {
			maxRight[j] = max(maxRight[j+1], nums[j])
		}
	}

	ans := make([]int, n-k+1)
	for i := 0; i+k <= n; i++ {
		ans[i] = max(maxRight[i], maxLeft[i+k-1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
