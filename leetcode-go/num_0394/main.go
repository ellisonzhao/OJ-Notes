package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	s := "abc3[cd]xyz"
	res := decodeString(s)
	fmt.Printf("res: %s", res)

	var cacheTime = 10 * time.Minute
	printInt(int(cacheTime.Seconds()))
}

func printInt(x int) {
	fmt.Printf("x int: %d", x)
}

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var resStack []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			// 字符出栈，保存中间结果
			var tempStack []byte
			for len(resStack) != 0 && resStack[len(resStack)-1] != '[' {
				val := resStack[len(resStack)-1]
				resStack = resStack[:len(resStack)-1]
				tempStack = append(tempStack, val)
			}
			// '[' 出栈
			if len(resStack) != 0 {
				resStack = resStack[:len(resStack)-1]
			}
			// 从后往前找数字起始位置
			count := 1
			for len(resStack) >= count &&
				resStack[len(resStack)-count] >= '0' &&
				resStack[len(resStack)-count] <= '9' {
				count++
			}
			// 注意索引边界问题
			digits := resStack[len(resStack)-count+1:]
			resStack = resStack[:len(resStack)-count+1]
			repTimes, _ := strconv.Atoi(string(digits))
			for j := 0; j < repTimes; j++ {
				// 把字符正向放回到栈里面
				for k := len(tempStack) - 1; k >= 0; k-- {
					resStack = append(resStack, tempStack[k])
				}
			}
		default:
			resStack = append(resStack, s[i])
		}
	}
	return string(resStack)
}
