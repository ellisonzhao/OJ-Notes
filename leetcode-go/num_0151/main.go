package num_0151

func reverseWords(s string) string {
	// 1. 使用双指针删除首尾空格
	var slow, fast int
	str := []byte(s)
	// 跳过头部冗余空格
	for len(str) > 0 && fast < len(str) && str[fast] == ' ' {
		fast++
	}
	// 处理单词间冗余空格
	for ; fast < len(str); fast++ {
		if fast-1 > 0 && str[fast-1] == str[fast] && str[fast] == ' ' {
			continue
		}
		str[slow] = str[fast]
		slow++
	}
	// 删除尾部冗余空格
	if slow-1 > 0 && str[slow-1] == ' ' {
		// 最后是空格
		str = str[:slow-1]
	} else {
		// 最后没有空格
		str = str[:slow]
	}
	// 2. 翻转整个字符串
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-1-i] = str[n-1-i], str[i]
	}
	// 3. 翻转单个单词

	for start := 0; start < n; {
		end := start
		for ; end < n && str[end] != ' '; end++ {
		}
		// 开始翻转
		for i, j := start, end-1; i < j; {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
		start = end + 1
	}
	return string(str)
}
