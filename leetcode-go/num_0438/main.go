package main

func findAnagrams(s string, p string) []int {
	window := make(map[int32]int)
	need := make(map[int32]int)
	for _, c := range p {
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	var res []int
	for right < len(s) {
		rightChar := int32(s[right])
		right++
		if need[rightChar] > 0 {
			window[rightChar]++
			if window[rightChar] == need[rightChar] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			leftChar := int32(s[left])
			left++
			if need[leftChar] > 0 {
				if window[leftChar] == need[leftChar] {
					valid--
				}
				window[leftChar]--
			}
		}
	}
	return res
}

func main() {

}
