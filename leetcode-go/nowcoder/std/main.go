package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func ReadInt(reader *bufio.Reader) int {
	num, _ := strconv.Atoi(ReadLine(reader))
	return num
}

func ReadArray(reader *bufio.Reader) []int {
	line := ReadLine(reader)
	strs := strings.Split(line, " ")
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	T := ReadInt(reader)
	for casi := int(0); casi < T; casi++ {
		n := ReadInt(reader)
		n = n / 2
		if n == 0 {
			n = 1
		}
		fmt.Println(n)
	}
}
