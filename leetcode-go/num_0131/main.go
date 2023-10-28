package main

import (
	"fmt"
	"reflect"
)

func partition(s string) [][]string {
	var result [][]string
	if s == "" || len(s) == 0 {
		return result
	}
	// 递归搜索函数
	var backtrack func(str string, paths []string)
	backtrack = func(str string, paths []string) {
		if str == "" || len(str) == 0 {
			result = append(result, append([]string(nil), paths...))
			return
		}
		// 遍历截取不同长度的子串
		for i := 1; i <= len(str); i++ {
			temp := str[0:i]
			if !isPalindrome(temp) {
				continue
			}
			paths = append(paths, temp)
			backtrack(str[i:], paths)
			paths = paths[:len(paths)-1]
		}
	}
	// 递归寻找结果
	backtrack(s, []string{})
	return result
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
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
