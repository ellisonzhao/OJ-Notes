package num_0093

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	if len(s) > 12 {
		return ans
	}
	var dfs func(restored string, idx, count int)
	dfs = func(restored string, idx, count int) {
		if count == 4 && idx == len(s) {
			ans = append(ans, restored)
			return
		}
		// 最多截取长度为 3
		for i := 1; i <= 3; i++ {
			if idx+i > len(s) {
				break
			}
			section := s[idx : idx+i]
			// 无效的前导 0
			if len(section) > 1 && section[0] == '0' {
				break
			}
			// 地址不能大于 255
			num, err := strconv.Atoi(section)
			if err != nil {
				panic(err)
			}
			if num > 255 {
				break
			}

			tmp := restored + section
			if count < 3 {
				tmp += "."
			}
			dfs(tmp, idx+i, count+1)
		}
	}

	dfs("", 0, 0)

	return ans
}
