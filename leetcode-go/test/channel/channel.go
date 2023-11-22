package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []int{1, 2, 3, 4, 5}
	results := ProcessTasks(tasks)

	count := 0
	for i := range results {
		fmt.Println(i)
		count++
	}

	fmt.Printf("Total processed tasks: %d\n", count)
}

func ProcessTasks(tasks []int) <-chan int {
	// your code
	results := make(chan int)
	var wg sync.WaitGroup
	go func() {
		defer close(results)
		for i := range tasks {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println(i)
				results <- i
			}(i)
		}
		wg.Wait()
	}()

	return results
}
