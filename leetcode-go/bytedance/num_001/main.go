package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

/**

数组 A 中给定可以使用的 1~9 的数，返回由 A 数组中的元素组成的小于 n 的最大数。
例如 A={1, 2, 4, 9}，n=2533，返回 2499
*/

func findMaxNumberLessThanN(arr []int, n int) int {
	sort.Ints(arr)

	maxNum := math.MinInt
	var dfs func(curNum int)
	dfs = func(curNum int) {
		// 剪枝：再往下遍历的数字只可能更大
		if len(strconv.Itoa(curNum)) == len(strconv.Itoa(n)) {
			return
		}
		for _, num := range arr {
			tmp := curNum*10 + num
			if tmp >= n {
				// 剪枝：返回上层
				return
			}
			maxNum = max(maxNum, tmp)
			dfs(tmp)
		}
	}

	dfs(0)

	if maxNum == math.MinInt {
		return 0
	}
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// getMaxDigitLtD 获取小于指定数字的数字。
func getMaxDigitLtD(digits []int, digit int) int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < digit {
			return digits[i]
		}
	}
	return 0
}

// getMaxNumLtN 获取小于 n 的最大数。
func getMaxNumLTN(nums []int, n int) int {
	var nDigits []int
	// 获取 n 的每一位数字，从低位到高位
	for n > 0 {
		nDigits = append(nDigits, n%10)
		n /= 10
	}

	// 排序给定的数字数组。
	sort.Ints(nums)

	// 数字写入 map 用于查看是否存在。
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	// 目标数的每一位数字。
	tDigits := make([]int, len(nDigits))

	// 从高位遍历，尽可能地取相同值，除了最后一位。
	for i := len(nDigits) - 1; i >= 0; i-- {
		if i > 0 {
			// 存在相同数字。
			if _, ok := m[nDigits[i]]; ok {
				tDigits[i] = nDigits[i]
				continue
			}
			// 存在小于当前位的最大数字。
			if d := getMaxDigitLtD(nums, nDigits[i]); d > 0 {
				tDigits[i] = d
				break
			}
		}
		if i == 0 {
			// 存在小于当前数字的最大数字。
			if d := getMaxDigitLtD(nums, nDigits[i]); d > 0 {
				tDigits[i] = d
				break
			}
		}
		// 数组中没有 <= 当前数字 的元素，进行回溯
		for ; i < len(nDigits); i++ {
			tDigits[i] = 0
			if d := getMaxDigitLtD(nums, nDigits[i]); d > 0 {
				tDigits[i] = d
				break
			}
			// 最高位也没有小于它的数字，结果取
			if i == len(nDigits)-1 {
				tDigits = tDigits[:len(tDigits)-1]
			}
		}
		// 找到了小于当前位的数字，后面全部取数组中最大数字
		break
	}

	// 拼接目标数。
	var target int
	for i := len(tDigits) - 1; i >= 0; i-- {
		target *= 10
		if tDigits[i] > 0 {
			target += tDigits[i]
		} else {
			target += nums[len(nums)-1]
		}
	}
	return target
}

func main() {
	// fmt.Println(findMaxNumberLessThanN([]int{4, 2, 5, 9}, 2451)) // 2449
	//
	// fmt.Println(getMaxNumLTN([]int{1, 9}, 991)) // 919
	//
	// fmt.Println(findMaxNumberLessThanN([]int{1, 2, 4, 9}, 2533)) // 2499

	fmt.Println(getMaxNumLTN([]int{2, 4, 9}, 1249)) // 999

	fmt.Println(findMaxNumberLessThanN([]int{2}, 1111)) // 222

	fmt.Println(findMaxNumberLessThanN([]int{2}, 33333)) // 22222

	fmt.Println(findMaxNumberLessThanN([]int{1, 2, 4, 9}, 1)) // 0

}
