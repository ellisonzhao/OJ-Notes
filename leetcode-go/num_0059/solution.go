package num_0059

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	left, right := 0, n-1
	top, bottom := 0, n-1
	val := 1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res[top][i] = val
			val++
		}
		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res[i][right] = val
			val++

		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res[bottom][i] = val
			val++

		}
		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			res[i][left] = val
			val++
		}
		left++
		if left > right {
			break
		}
	}

	return res
}
