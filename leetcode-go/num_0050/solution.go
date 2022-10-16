package num_0050

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	res := float64(1)
	for n > 0 {
		if (n & 1) == 1 {
			res *= x
		}
		x *= x
		n = n >> 1
	}
	return res
}
