package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// 统计目标字符串中的字符出现次数
	need := make(map[int]int)
	// 统计「滑动窗口内」目标字符串中的字符出现次数
	window := make(map[int]int)
	for _, c := range t {
		need[int(c)]++
	}

	left, right := 0, 0
	start, width := left, math.MaxInt32
	valid := 0
	for right < len(s) {
		// 右侧移入窗口内的字符
		rightChar := s[right]
		// 窗口扩大
		right++
		// 添加到窗口内的字符是目标字符，窗口内数据更新
		if need[int(rightChar)] > 0 {
			window[int(rightChar)]++
			// 窗口内该目标字符出现次数已满足条件
			if window[int(rightChar)] == need[int(rightChar)] {
				valid++
			}
		}

		// 满足条件时，更新可能的结果并收缩窗口
		// 此时窗口内已经包含目标字符串中的所有字符
		for valid == len(need) {
			// 更新最小覆盖子串
			if right-left < width {
				start = left
				width = right - left
			}
			// 即将从窗口左侧移除的字符
			leftChar := s[left]
			// 窗口缩小
			left++
			// 即将从窗口内移除的是目标字符，窗口内数据更新
			if need[int(leftChar)] > 0 {
				if window[int(leftChar)] == need[int(leftChar)] {
					// 窗口内该目标字符出现次数将不满足条件
					valid--
				}
				window[int(leftChar)]--
			}
		}
	}

	if width == math.MaxInt32 {
		return ""
	}
	return s[start : start+width]
}

// func minWindow(s string, t string) string {
// 	if len(s) < len(t) {
// 		return ""
// 	}
//
// 	freq := make(map[int]int)
// 	for _, c := range t {
// 		freq[int(c)]++
// 	}
// 	count := len(t)
//
// 	left, right := 0, 0
// 	start := left
// 	width := math.MaxInt32
// 	for right < len(s) {
// 		rightChar := s[right]
// 		right++
// 		// 窗口内数据更新
// 		if freq[int(rightChar)] > 0 {
// 			count--
// 		}
// 		freq[int(rightChar)]--
//
// 		// 满足条件时，更新可能的结果并收缩窗口
// 		for count == 0 {
// 			if right-left < width {
// 				start = left
// 				width = right - left
// 			}
//
// 			leftChar := s[left]
// 			// 如果是t中出现的字符，此时 freq[int(leftChar)] = 0
// 			freq[int(leftChar)]++
// 			if freq[int(leftChar)] > 0 {
// 				count++
// 			}
// 			left++
// 		}
// 	}
//
// 	if width == math.MaxInt32 {
// 		return ""
// 	}
// 	return s[start : start+width]
// }

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Println(minWindow(s, t))
}
