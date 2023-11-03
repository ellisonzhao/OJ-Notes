package num_0001

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		rem := target - num
		if j, ok := seen[rem]; ok {
			return []int{i, j}
		}
		seen[num] = i
	}

	return []int{}
}
