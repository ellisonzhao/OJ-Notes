package num_0003

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	charIndexes := make(map[uint8]int)
	result := 0
	for start, end := 0, 0; end < len(s); end++ {
		// 窗口右侧字符已存在
		if idx, ok := charIndexes[s[end]]; ok {
			start = max(start, idx+1)
		}

		charIndexes[s[end]] = end

		result = max(result, end-start+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
