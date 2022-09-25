package num_0003

func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int)
	left, right := 0, 0
	length, res := 0, 0
	for right < len(s) {
		rc := s[right]
		if window[rc] > 0 && window[rc] >= left {
			left = window[rc] + 1
			length = right - left
		}

		window[rc] = right
		length++
		right++
		if length > res {
			res = length
		}
	}
	return res
}
