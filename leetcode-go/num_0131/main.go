package main

import (
	"fmt"
	"reflect"
)

//	func partition(s string) [][]string {
//		var result [][]string
//		if s == "" || len(s) == 0 {
//			return result
//		}
//		// 递归搜索函数
//		var backtrack func(str string, paths []string)
//		backtrack = func(str string, paths []string) {
//			if str == "" || len(str) == 0 {
//				result = append(result, append([]string(nil), paths...))
//				return
//			}
//			// 遍历截取不同长度的子串
//			for i := 1; i <= len(str); i++ {
//				temp := str[0:i]
//				if !isPalindrome(temp) {
//					continue
//				}
//				paths = append(paths, temp)
//				backtrack(str[i:], paths)
//				paths = paths[:len(paths)-1]
//			}
//		}
//		// 递归寻找结果
//		backtrack(s, []string{})
//		return result
//	}
//
//	func isPalindrome(s string) bool {
//		left, right := 0, len(s)-1
//		for left < right {
//			if s[left] != s[right] {
//				return false
//			}
//			left++
//			right--
//		}
//		return true
//	}

func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	// 记录回文子串
	for k := 0; k < n; k++ {
		for j := 0; j <= k; j++ {
			if s[j] != s[k] {
				continue
			}
			if k-j <= 2 {
				// a、aa、aba 格式
				dp[j][k] = true
			} else {
				dp[j][k] = dp[j+1][k-1]
			}
		}
	}
	// 递归处理
	var result [][]string
	var backtrack func(paths []string, s string, pos int)
	backtrack = func(paths []string, s string, pos int) {
		if pos == len(s) {
			result = append(result, append([]string(nil), paths...))
			return
		}
		for i := pos; i < len(s); i++ {
			if dp[pos][i] {
				paths = append(paths, s[pos:i+1])
				backtrack(paths, s, i+1)
				paths = paths[:len(paths)-1]
			}
		}
	}
	backtrack([]string{}, s, 0)
	return result
}

func main() {
	s := "aab"
	for i, ch := range s {
		fmt.Println(i)
		fmt.Println(reflect.TypeOf(s[i]))
		fmt.Println(reflect.TypeOf(ch))
	}
	fmt.Println(partition(s))
}
