package num_0402

import (
	"strings"
)

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := range num {
		digit := num[i]
		// 模拟删除 k 个数字
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	// 如果 k>0 说明还要删掉剩余的 k 个数字
	stack = stack[:len(stack)-k]
	// 去掉前导 0
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
