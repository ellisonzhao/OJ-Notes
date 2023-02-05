package num_0316

func removeDuplicateLetters(s string) string {
	var counter [128]int
	for _, ch := range s {
		// type of ch: int32
		counter[ch]++
	}

	var (
		// type of byte: uint8
		stack   []byte
		inStack [128]bool
	)
	for i := range s {
		// type of s[i]: uint8
		ch := s[i]
		if inStack[ch] {
			counter[ch]--
			continue
		}
		// 从栈底到栈顶，字母的字典序只增不减
		for len(stack) > 0 && stack[len(stack)-1] > ch {
			last := stack[len(stack)-1]
			// 剩余字符串中不会再次出现栈顶字符，保留
			if counter[last] == 0 {
				break
			}
			// 栈顶字母不是最后一次出现，出栈
			stack = stack[:len(stack)-1]
			inStack[last] = false
		}
		stack = append(stack, ch)
		inStack[ch] = true
		counter[ch]--
	}

	return string(stack)
}
