package main

func main() {

}

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i, rate := range ratings {
		if i > 0 && rate > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}

	right := 0
	ans := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(candies[i], right)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
