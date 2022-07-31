package main

import "strconv"

func main() {

}

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		// 数字直接进栈
		if err == nil {
			stack = append(stack, val)
			continue
		}
		// 遇到运算符，数字出栈。题目中表达式总是合法的，因此这里也没有对栈进行判空
		numFirst, numSecond := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		switch token {
		case "+":
			stack = append(stack, numFirst+numSecond)
		case "-":
			stack = append(stack, numFirst-numSecond)
		case "*":
			stack = append(stack, numFirst*numSecond)
		default:
			stack = append(stack, numFirst/numSecond)
		}
	}
	return stack[0]
}
