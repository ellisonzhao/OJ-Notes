package num_0402

import (
	"strings"
)

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return ""
	}

	var stack []byte
	for i := range num {
		digit := num[i]
		// 例如栈内是 14，此时遇到 3，显然 13 是比 14 更小的数，因此更新栈顶元素
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		// 维护单调不减的数字序列
		stack = append(stack, digit)
	}
	// 数字序列已经是升序，取前 n-k 个构成的结果最小，即去掉尾部 k 个数字
	stack = stack[:len(stack)-k]
	// 去掉前导 0
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
