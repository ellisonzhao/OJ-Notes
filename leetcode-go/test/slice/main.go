package main

import (
	"fmt"
	"sync"
)

func main() {
	//
	var wg sync.WaitGroup
	var arr []int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				arr = append(arr, 1)
			}()
		}
	}

	wg.Wait()

	fmt.Println(len(arr))
}
