package main

import (
	"fmt"
)

// func subsets(nums []int) [][]int {
// 	var ans [][]int
// 	n := len(nums)
// 	for mask := 0; mask < (1 << n); mask++ {
// 		var subset []int
// 		for i, v := range nums {
// 			// 二进制位为 1 表示「在子集中」
// 			if ((mask >> i) & 1) > 0 {
// 				subset = append(subset, v)
// 			}
// 		}
// 		ans = append(ans, subset)
// 	}
// 	return ans
// }

// func subsets(nums []int) [][]int {
// 	var (
// 		ans [][]int
// 		dfs func(int, []int)
// 	)
// 	dfs = func(idx int, tmp []int) {
// 		if idx == len(nums) {
// 			ans = append(ans, append([]int(nil), tmp...))
// 			return
// 		}
// 		dfs(idx+1, tmp)
// 		tmp = append(tmp, nums[idx])
// 		dfs(idx+1, tmp)
// 		tmp = tmp[:len(tmp)-1]
// 	}
// 	dfs(0, []int(nil))
// 	return ans
// }

func subsets(nums []int) [][]int {
	var (
		ans [][]int
		dfs func(int, []int)
	)
	dfs = func(idx int, tmp []int) {
		ans = append(ans, append([]int(nil), tmp...))
		for i := idx; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			dfs(i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(0, []int{})

	return ans
}

func main() {
	nums := []int{9, 0, 3, 5, 7}
	ans := subsets(nums)
	fmt.Println(ans)
}
