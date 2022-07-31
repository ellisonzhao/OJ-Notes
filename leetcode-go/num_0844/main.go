package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	skipS, skipT := 0, 0

	i, j := len(S)-1, len(T)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if S[i] == '#' {
				// 退格字符数量增加
				skipS++
				i--
			} else if skipS > 0 {
				// 根据退格数量跳过退格前的字符
				skipS--
				i--
			} else {
				break
			}
		}

		for j >= 0 {
			if T[j] == '#' {
				// 退格字符数量增加
				skipT++
				j--
			} else if skipT > 0 {
				// 根据退格数量跳过退格前的字符
				skipT--
				j--
			} else {
				break
			}
		}

		// 不相等的情形
		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				// 当前有效位置字符不相等
				return false
			}
		} else if i >= 0 || j >= 0 {
			// 退格后字符串长度不相等
			return false
		}

		i--
		j--
	}

	return true
}

func main() {
	s, t := "a#c", "b"
	fmt.Println(backspaceCompare(s, t))
}
