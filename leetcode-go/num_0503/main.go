package num_0503

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var stack []int
	for i := 0; i < 2*n-1; i++ {
		// 当前遍历位置对应数组中下标
		idx := i % n
		// 维护单调递减栈，栈内存放元素下标
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[idx] {
			// 找到目标结果，出栈
			ans[stack[len(stack)-1]] = nums[idx]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, idx)
	}
	// 遍历结束后，栈内元素是没有找到目标值的元素下标

	return ans
}
