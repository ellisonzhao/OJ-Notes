package num_0054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	row, col := len(matrix), len(matrix[0])
	res := make([]int, 0, row*col)
	left, right := 0, col-1
	top, bottom := 0, row-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}
