package num_0039

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	// 定义递归函数，避免全局变量
	var dfs func(nums []int, remain, start int)
	dfs = func(nums []int, remain, start int) {
		if remain < 0 {
			return
		}
		// 已选元素之和等于 target
		if remain == 0 {
			ans = append(ans, append([]int(nil), nums...))
			return
		}
		// 递归搜索所有可能的结果
		for i := start; i < len(candidates); i++ {
			num := candidates[i]
			nums = append(nums, num)
			// 可重复使用当前元素
			dfs(nums, remain-num, i)
			nums = nums[:len(nums)-1]
		}
	}

	dfs([]int(nil), target, 0)
	return ans
}
