package num_0056

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// intervals.length >= 1
	// intervals[i].length == 2
	sort.Slice(intervals, func(i, j int) bool {
		// [start, end] 按 start 升序排序
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		// 与前一个区间没有交集
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
			continue
		}
		// 要与前一个区间合并，判断是否更新区间右边界
		if intervals[i][1] > end {
			end = intervals[i][1]
		}
	}
	// 最后剩下的区间
	ans = append(ans, []int{start, end})
	return ans
}
