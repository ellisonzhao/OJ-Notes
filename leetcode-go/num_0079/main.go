package num_0079

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for i := range board {
		for j := range board[i] {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j int, word string, pos int) bool {
	if pos == len(word) {
		return true
	}
	// 越界或字符不能组成单词
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != word[pos] {
		return false
	}
	// 暂存当前位置字符
	curr := board[i][j]
	// 标记为已访问过
	board[i][j] = '#'
	// 继续回溯下一个字符
	res := dfs(board, i+1, j, word, pos+1) ||
		dfs(board, i, j+1, word, pos+1) ||
		dfs(board, i-1, j, word, pos+1) ||
		dfs(board, i, j-1, word, pos+1)
	// 恢复当前位置字符
	board[i][j] = curr
	return res
}
