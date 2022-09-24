package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 打开读取文件
	infile, err := os.Open("/Users/ellison/go/src/lc-solution-go/solution/num_2/input.data")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	stats, err := infile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// 读取字节
	bytes := make([]byte, stats.Size())
	buf := bufio.NewReader(infile)
	_, err = buf.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	// 统计出现次数
	counter := make(map[byte]int)
	var maxCount int
	for _, b := range bytes {
		_, ok := counter[b]
		if !ok {
			counter[b] = 1
			continue
		}

		counter[b]++
		if counter[b] > maxCount {
			maxCount = counter[b]
		}
	}

	// 输出数值和出现次数
	for k, v := range counter {
		if v == maxCount {
			fmt.Println(k, v)
		}
	}
}
