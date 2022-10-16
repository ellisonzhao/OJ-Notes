package num_0033

func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[0] {
			// 左半部分有序区间
			if nums[0] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			// 右半部分有序区间
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
