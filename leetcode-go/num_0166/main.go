package num_0166

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var res []byte
	// 异号结果为负
	if (numerator > 0) != (denominator > 0) {
		res = append(res, '-')
	}
	num := abs(numerator)
	den := abs(denominator)
	// 整数部分
	integer := num / den
	res = append(res, strconv.Itoa(integer)...)
	num %= den
	if num == 0 {
		return string(res)
	}
	// 小数部分
	res = append(res, '.')
	remainderIndexes := map[int]int{}
	remainderIndexes[num] = len(res)
	for num != 0 {
		num *= 10
		// 数字转 byte
		res = append(res, strconv.Itoa(num/den)...)
		num %= den
		// 存在循环小数
		if _, ok := remainderIndexes[num]; ok {
			insertIndex := remainderIndexes[num]
			// 取循环部分
			res = append(res[:insertIndex], append([]byte{'('}, res[insertIndex:]...)...)
			res = append(res, ')')
			break
		} else {
			remainderIndexes[num] = len(res)
		}
	}
	return string(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
