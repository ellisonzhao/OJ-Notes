package main

import (
	"fmt"
	"math"
)

func commonChars(A []string) (res []string) {
	freq := [26]int{}
	for i := range freq {
		freq[i] = math.MaxInt64
	}

	for _, word := range A {
		counter := [26]int{}
		// 统计该单词中各个字符的出现次数
		for _, c := range word {
			counter[c-'a']++
		}
		// 保存字符出现的最少次数
		for i, f := range counter {
			if f < freq[i] {
				freq[i] = f
			}
		}
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < freq[i]; j++ {
			res = append(res, fmt.Sprint('a'+i))
		}
	}

	return
}

func main() {
	A := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(A))
}
