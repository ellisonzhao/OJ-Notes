package num_0040

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
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
			// 有多个相同元素时，只添加第一个，后续跳过重复元素
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			num := candidates[i]
			nums = append(nums, num)
			// 每个数字在每个组合中只能使用 1 次
			dfs(nums, remain-num, i+1)
			nums = nums[:len(nums)-1]
		}
	}

	dfs([]int(nil), target, 0)

	return ans
}
