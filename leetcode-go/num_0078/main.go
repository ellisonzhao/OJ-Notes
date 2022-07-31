package main

// func subsets(nums []int) (ans [][]int) {
//	n := len(nums)
//	for mask := 0; mask < (1 << n); mask++ {
//		var set []int
//		for i, v := range nums {
//			if ((mask >> i) & 1) > 0 {
//				set = append(set, v)
//			}
//		}
//		// tmp := []int{}
//		ans = append(ans, append([]int{}, set...))
//	}
//	return
// }

// func subsets(nums []int) (ans [][]int) {
//	var set []int
//	var dfs func(int)
//	dfs = func(cur int) {
//		if cur == len(nums) {
//			ans = append(ans, append([]int(nil), set...))
//			return
//		}
//		set = append(set, nums[cur])
//		dfs(cur + 1)
//		set = set[:len(set)-1]
//		dfs(cur + 1)
//	}
//	dfs(0)
//	return
// }

func subsets(nums []int) (results [][]int) {
	var temp []int
	var dfs func(int, []int)
	dfs = func(idx int, temp []int) {
		if idx == len(nums)+1 {
			return
		}

		results = append(results, append([]int{}, temp...))

		for i := idx; i < len(nums); i++ {
			temp = append(temp, nums[idx])
			dfs(i+1, temp)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0, temp)

	return
}
