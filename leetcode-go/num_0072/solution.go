package main

import "fmt"

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 {
		return l2
	}
	if l2 == 0 {
		return l1
	}
	// 初始化 dp 数组，word1 前 i 个字符转换成 word2 前 j 个字符所使用的最少操作数
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	for j := 1; j <= l2; j++ {
		dp[0][j] = j
	}
	// 从头开始计算
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 插入一个字符
				insertOp := 1 + dp[i][j-1]
				// 删除一个字符
				deleteOp := 1 + dp[i-1][j]
				// 替换一个字符
				replaceOp := 1 + dp[i-1][j-1]
				dp[i][j] = min(insertOp, min(deleteOp, replaceOp))
			}
		}
	}
	return dp[l1][l2]

}

// func minDistance(word1 string, word2 string) int {
// 	l1, l2 := len(word1), len(word2)
// 	if l1 == 0 {
// 		return l2
// 	}
// 	if l2 == 0 {
// 		return l1
// 	}
// 	// 初始化 dp 数组，word1 前 i 个字符转换成 word2 前 j 个字符所使用的最少操作数
// 	dp := make([]int, l2+1)
// 	for i := 0; i <= l2; i++ {
// 		dp[i] = i
// 	}
// 	// 从头开始计算
// 	for i := 1; i <= l1; i++ {
// 		pre := dp[0]
// 		dp[0] = i
// 		for j := 1; j <= l2; j++ {
// 			tmp := dp[j]
// 			if word1[i-1] == word2[j-1] {
// 				dp[j] = pre
// 			} else {
// 				// 插入一个字符
// 				insertOp := 1 + dp[j]
// 				// 删除一个字符
// 				deleteOp := 1 + dp[j-1]
// 				// 替换一个字符
// 				replaceOp := 1 + pre
// 				dp[j] = min(insertOp, min(deleteOp, replaceOp))
// 			}
// 			pre = tmp
// 		}
// 	}
// 	return dp[l2]
//
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
