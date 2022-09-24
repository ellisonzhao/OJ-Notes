package main

import "container/list"

func main() {

}

func calculate(s string) int {
	stack := list.New()
	// 记录加减状态
	sign := 1
	result := 0
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		} else if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] <= '9' && s[i] >= '0' { // 获取一段数字
			num := 0
			for i < len(s) && s[i] <= '9' && s[i] >= '0' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += num * sign
		} else if s[i] == '(' { // 把之前计算结果及加减状态压栈，开始新的计算
			stack.PushBack(result)
			stack.PushBack(sign)
			result = 0
			sign = 1
			i++
		} else if s[i] == ')' { // 新的计算结果 * 前一个加减状态 + 之前计算结果
			result = result*stack.Remove(stack.Back()).(int) + stack.Remove(stack.Back()).(int)
			i++
		}
	}
	return result
}
