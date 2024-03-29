package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	pieces1 := strings.Split(version1, ".")
	pieces2 := strings.Split(version2, ".")
	n := len(pieces1)
	if len(pieces2) > n {
		n = len(pieces2)
	}

	for i := 0; i < n; i++ {
		var num1, num2 int
		if i >= len(pieces1) {
			num1 = 0
		} else {
			num1, _ = strconv.Atoi(pieces1[i])
		}
		if i >= len(pieces2) {
			num2 = 0
		} else {
			num2, _ = strconv.Atoi(pieces2[i])
		}
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}

	return 0
}

func strToInt(str string) int {
	var num int
	for i := 0; i < len(str); i++ {
		num = num*10 + int(str[i]-'0')
	}
	return num
}

func compareVersion1(version1 string, version2 string) int {
	var num1, num2 int
	len1, len2 := len(version1), len(version2)
	for i, j := 0, 0; i < len1 || j < len2; {
		num1, num2 = 0, 0
		for i < len1 && version1[i] != '.' {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}
		for j < len2 && version2[j] != '.' {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}

		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		} else {
			i++
			j++
		}
	}

	return 0
}

func main() {
	version1 := "1.01"
	version2 := "1.001"
	res := compareVersion(version1, version2)
	fmt.Println(res)
}
