package num_0005

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 0 {
		return ""
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	res := ""
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = (s[i] == s[j]) && (j-i <= 2 || dp[i+1][j-1])

			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
