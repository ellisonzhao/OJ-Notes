package main

func main() {

}

func canJump(nums []int) bool {
	n := len(nums)
	maxIndex := 0
	for i := 0; i < n; i++ {
		if i <= maxIndex {
			if maxIndex = max(maxIndex, i+nums[i]); maxIndex >= n-1 {
				return true
			}
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
