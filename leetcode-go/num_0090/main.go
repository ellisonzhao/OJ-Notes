package num_0090

import (
	"sort"
)

// func subsetsWithDup(nums []int) [][]int {
// 	sort.Ints(nums)
// 	var (
// 		ans [][]int
// 		dfs func(idx int, tmp []int)
// 	)
// 	dfs = func(idx int, tmp []int) {
// 		// 遇到相同的元素会先添加第一个
// 		ans = append(ans, append([]int(nil), tmp...))
// 		for i := idx; i < len(nums); i++ {
// 			if i > idx && nums[i] == nums[i-1] {
// 				continue
// 			}
// 			tmp = append(tmp, nums[i])
// 			dfs(i+1, tmp)
// 			tmp = tmp[:len(tmp)-1]
// 		}
// 	}
// 	dfs(0, []int(nil))
// 	return ans
// }

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var (
		ans [][]int
		dfs func(subset []int, idx int)
	)
	dfs = func(subset []int, idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int(nil), subset...))
			return
		}
		// 添加当前元素
		subset = append(subset, nums[idx])
		dfs(subset, idx+1)
		// 不添加当前元素，需要跳过重复
		subset = subset[:len(subset)-1]
		for idx < len(nums)-1 && nums[idx] == nums[idx+1] {
			idx++
		}
		dfs(subset, idx+1)
	}
	dfs([]int(nil), 0)

	return ans
}
