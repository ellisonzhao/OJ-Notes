package num_0022

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(curr string, lCount, rCount int)
	dfs = func(curr string, lCount, rCount int) {
		if lCount == n && rCount == n {
			ans = append(ans, curr)
			return
		}
		// 左括号数量最多为 n
		// 当前字符串可以添加左括号
		if lCount < n {
			dfs(curr+"(", lCount+1, rCount)
		}
		// 右括号数量最多为 n
		// 已添加的左括号数量必须大于已添加的右括号数量，才可以添加右括号
		if lCount > rCount {
			dfs(curr+")", lCount, rCount+1)
		}
	}

	dfs("", 0, 0)

	return ans
}
