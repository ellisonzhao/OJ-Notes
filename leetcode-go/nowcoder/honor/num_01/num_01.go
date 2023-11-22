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

func main() {
	fmt.Scan()
	reader := bufio.NewReader(os.Stdin)
	n := ReadInt(reader)
	res := helper(n)
	fmt.Println(res)
}

func helper(n int) int64 {
	pre := 1
	var sum int64
	for i := 1; i <= n; i++ {
		pre *= i
		sum += int64(pre)
	}
	return sum
}
