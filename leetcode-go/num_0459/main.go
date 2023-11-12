package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n < 1 {
		return false
	}
	s1 := s + s
	s1 = s1[1 : 2*n-1]
	if strings.Contains(s1, s) {
		return true
	}
	return false
}

func main() {

}
