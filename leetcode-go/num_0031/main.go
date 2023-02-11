package num_0031

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 对于形如 [3,2,1] 降序的序列，i = -1，直接逆置即可
	if i < 0 {
		reverse(nums[i+1:])
		return
	}
	// 从右往左找到第一个 nums[j] > nums[i] 的位置
	j := len(nums) - 1
	for j >= 0 && nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
