package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[int]int{1: 1}
	go doWrite(m)
	go doWrite(m)

	time.Sleep(1 * time.Second)
	fmt.Println(m)
}

func doWrite(m map[int]int) {
	i := 0
	for i < 10000 {
		m[1] = 1
		i++
	}
}

func doRead(m map[int]int) {
	i := 0
	for i < 10000 {
		if _, ok := m[1]; ok {
			i++
		}
	}
}
