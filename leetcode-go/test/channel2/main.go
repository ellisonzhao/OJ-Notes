package main

import (
	"fmt"
	"sync"
)

var POOL = 100

func groutine1(first chan struct{}, second chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= POOL; i++ {
		<-first
		fmt.Print("A")
		second <- struct{}{}
	}

	<-first
}

func groutine2(second chan struct{}, third chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= POOL; i++ {
		<-second
		fmt.Print("B")
		third <- struct{}{}

	}
}

func groutine3(third chan struct{}, first chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= POOL; i++ {
		<-third
		fmt.Println("C")
		first <- struct{}{}
	}
}

func main() {
	first := make(chan struct{})
	second := make(chan struct{})
	third := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go groutine1(first, second, wg)
	go groutine2(second, third, wg)
	go groutine3(third, first, wg)

	first <- struct{}{}
	wg.Wait()

	close(first)
	close(second)
	close(third)

}

// package main
//
// import (
// 	"fmt"
// 	"sync"
// )
//
// func main() {
// 	var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go func(s string) {
// 		defer wg.Done()
// 		for i := 1; i <= 10; i++ {
// 			<-ch1
// 			fmt.Print(s)
// 			ch2 <- struct{}{}
// 		}
// 		<-ch1
// 	}("A")
// 	go func(s string) {
// 		defer wg.Done()
// 		for i := 1; i <= 10; i++ {
// 			<-ch2
// 			fmt.Print(s)
// 			ch3 <- struct{}{}
// 		}
// 	}("B")
// 	go func(s string) {
// 		defer wg.Done()
// 		for i := 1; i <= 10; i++ {
// 			<-ch3
// 			fmt.Println(s)
// 			ch1 <- struct{}{}
// 		}
// 	}("C")
// 	ch1 <- struct{}{}
// 	wg.Wait()
//
// 	close(ch1)
// 	close(ch2)
// 	close(ch3)
//
// }
