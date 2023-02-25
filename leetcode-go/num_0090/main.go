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
		tmp []int
		ans [][]int
		dfs func(idx int, choosePre bool)
	)
	dfs = func(idx int, choosePre bool) {
		if idx == len(nums) {
			ans = append(ans, append([]int(nil), tmp...))
			return
		}
		dfs(idx+1, false)
		if !choosePre && idx > 0 && nums[idx-1] == nums[idx] {
			return
		}
		tmp = append(tmp, nums[idx])
		dfs(idx+1, true)
		tmp = tmp[:len(tmp)-1]
	}
	dfs(0, false)
	return ans
}
