package num_0046

import (
	"sort"
)

var result [][]int

func permuteUnique(nums []int) [][]int {
	result = make([][]int, 0, len(nums))
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	sort.Ints(nums)
	backtrack(nums, path, used)

	return result
}

func backtrack(nums, path []int, used []bool) {
	if len(path) == len(nums) {
		p := make([]int, len(path))
		copy(p, path)

		result = append(result, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		used[i] = true
		path = append(path, nums[i])

		backtrack(nums, path, used)

		used[i] = false
		path = path[:len(path)-1]
	}
}
