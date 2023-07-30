package num_0008

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	idx, sign := 0, 1
	if s[idx] == '-' || s[idx] == '+' {
		if s[idx] == '-' {
			sign = -1
		}
		idx++
	}

	var base int
	for ; idx < len(s) && s[idx] >= '0' && s[idx] <= '9'; idx++ {
		base = base*10 + int(s[idx]-'0')
		if base*sign > math.MaxInt32 {
			return math.MaxInt32
		}
		if base*sign < math.MinInt32 {
			return math.MinInt32
		}
	}
	base *= sign

	return base
}
