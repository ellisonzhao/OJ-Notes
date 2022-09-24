package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AppendTail(5)
	obj.AppendTail(2)
	fmt.Printf("stack1: %v", obj.FirstStack)
	obj.DeleteHead()
	obj.DeleteHead()

}

func jump(nums []int) int {
	n := len(nums)
	f := make([]int, len(nums))
	f[0] = 0

	for i := 1; i < n; i++ {
		f[i] = i
		// 遍历之前结果取一个最小值+1
		for j := 0; j < i; j++ {
			// 位置 j 能够到达 i
			if nums[j]+j >= i {
				f[i] = min(f[j]+1, f[i])
			}
		}
	}

	return f[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type CQueue struct {
	FirstStack  []int
	SecondStack []int
}

func Constructor() CQueue {
	firstStack := make([]int, 0)
	secondStack := make([]int, 0)
	return CQueue{
		FirstStack:  firstStack,
		SecondStack: secondStack,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.FirstStack = append(this.FirstStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.SecondStack) == 0 {
		n := len(this.FirstStack)
		for i := n - 1; i >= 0; i-- {
			this.SecondStack = append(this.SecondStack, this.FirstStack[i])
			this.FirstStack = this.FirstStack[:i]
		}
	}

	fmt.Printf("stack2: %v", this.SecondStack)

	if n := len(this.SecondStack); n != 0 {
		res := this.SecondStack[n-1]
		this.SecondStack = this.SecondStack[:n-1]

		return res
	}

	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
