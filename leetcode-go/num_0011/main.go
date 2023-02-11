package num_0011

func maxArea(height []int) int {
	ans := 0
	i, j := 0, len(height)-1
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		ans = max(ans, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
