package main

import (
	"fmt"
)

func main() {
	m := 0
	n := 0
	for {
		t, _ := fmt.Scan(&m, &n)
		if t == 0 {
			break
		}
		// 图顶点记录
		graph := make([][]int, m+1)
		for i := 0; i <= m; i++ {
			graph[i] = make([]int, m+1)
		}
		// 记录 n 条线路以及距离
		for i := 0; i < n; i++ {
			var a, b, c, d int
			t, _ = fmt.Scan(&a, &b, &c, &d)
			if t == 0 {
				continue
			}
			graph[b][c] = d
			graph[c][b] = d
		}
		// 计算顶点之间的最短距离
		var u, v int
		t, _ = fmt.Scan(&u, &v)
		if t == 0 {
			break
		}

		fmt.Println(graph)
	}
}

// 迪杰斯特拉算法计算最短路径
func helper(graph [][]int, u, v int) int {
	//
	return 0
}
