package num_0496

// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	greaterNums := make(map[int]int)
// 	var stack []int
// 	for _, num := range nums2 {
// 		// 维护单调递减栈
// 		for len(stack) > 0 && stack[len(stack)-1] < num {
// 			// 记录栈顶元素的下一个更大元素
// 			greaterNums[stack[len(stack)-1]] = num
// 			// 出栈
// 			stack = stack[:len(stack)-1]
// 		}
// 		stack = append(stack, num)
// 	}
// 	// 遍历结束后，栈内元素是没有找到目标值的元素
// 	// 找到目标值的元素已经在 greaterNums 中
// 	res := make([]int, 0, len(nums1))
// 	for _, num := range nums1 {
// 		if greater, ok := greaterNums[num]; ok {
// 			res = append(res, greater)
// 		} else {
// 			res = append(res, -1)
// 		}
// 	}
//
// 	return res
// }

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greaterNums := make(map[int]int)
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		// 维护一个单调递减栈：栈底到栈顶元素依次减小
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		// 栈顶元素就是第一个从左到右更大的元素，用 map 保存
		if len(stack) > 0 {
			greaterNums[num] = stack[len(stack)-1]
		} else {
			// 不存在
			greaterNums[num] = -1
		}
		stack = append(stack, num)
	}
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = greaterNums[num]
	}
	return res
}
