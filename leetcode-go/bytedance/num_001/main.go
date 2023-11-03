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

func main() {
	fmt.Println(findMaxNumberLessThanN([]int{4, 2, 5, 9}, 2451)) // 2449

	fmt.Println(findMaxNumberLessThanN([]int{1, 9}, 991)) // 919

	fmt.Println(findMaxNumberLessThanN([]int{1, 2, 4, 9}, 2533)) // 2499

	fmt.Println(findMaxNumberLessThanN([]int{2}, 1111)) // 222

	fmt.Println(findMaxNumberLessThanN([]int{2}, 33333)) // 22222

	fmt.Println(findMaxNumberLessThanN([]int{1, 2, 4, 9}, 1)) // 22222

}
