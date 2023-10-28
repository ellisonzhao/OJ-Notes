package main

func main() {

}

func canJump(nums []int) bool {
	var farthestIdx int
	for i, num := range nums {
		if i <= farthestIdx {
			// 遍历每个下标，找能跳到的最远位置
			farthestIdx = max(farthestIdx, i+num)
			if farthestIdx >= len(nums)-1 {
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
