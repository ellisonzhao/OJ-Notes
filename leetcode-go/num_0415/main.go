package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var ans string
	carry := 0
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum := x + y + carry
		ans = strconv.Itoa(sum%10) + ans
		carry = sum / 10
		i--
		j--
	}
	if carry > 0 {
		ans = strconv.Itoa(carry) + ans
	}
	return ans
}

func main() {

}
