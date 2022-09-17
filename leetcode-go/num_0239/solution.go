package main

import (
	"container/heap"
	"sort"
)

var arr []int

type maxHeap struct{ sort.IntSlice }

func (h maxHeap) Less(i, j int) bool  { return arr[h.IntSlice[i]] > arr[h.IntSlice[j]] }
func (h *maxHeap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *maxHeap) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	arr = nums
	q := &maxHeap{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
