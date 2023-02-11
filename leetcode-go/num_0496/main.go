package num_0496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greaterMap := make(map[int]int)
	var stack []int
	for _, num := range nums2 {
		// 维护单调递减栈
		for len(stack) > 0 && stack[len(stack)-1] < num {
			// 记录栈顶元素的目标值
			greaterMap[stack[len(stack)-1]] = num
			// 出栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	// 遍历结束后，栈内元素是没有找到目标值的元素
	// 找到目标值的元素已经在 greaterMap 中
	res := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		if greater, ok := greaterMap[num]; ok {
			res = append(res, greater)
		} else {
			res = append(res, -1)
		}
	}

	return res
}
