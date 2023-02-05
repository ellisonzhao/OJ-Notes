package num_0128

func longestConsecutive(nums []int) int {
	var res int
	numsMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			continue
		}
		// num 不在 map 中，获取 num-1 和 num+1 所在连续区间长度
		leftLen, rightLen := 0, 0
		// left 为 num-1 所在连续区间的长度，更进一步理解为：左连续区间的长度
		if left, ok := numsMap[num-1]; ok {
			leftLen = left
		}
		// right 为 num+1 所在连续区间的长度，更进一步理解为：右连续区间的长度
		if right, ok := numsMap[num+1]; ok {
			rightLen = right
		}
		// 当前连续区间的总长度
		length := leftLen + rightLen + 1
		if length > res {
			res = length
		}
		// 更新当前连续区间左边界和右边界对应的区间长度
		numsMap[num] = length
		numsMap[num-leftLen] = length
		numsMap[num+rightLen] = length
	}

	return res
}
