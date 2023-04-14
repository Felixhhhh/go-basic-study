package main

//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	test()
//	ch1 := make(chan int, 1)
//	ch2 := make(chan int, 1)
//	ch3 := make(chan int, 1)
//	ch1 <- 1
//
//	var wg sync.WaitGroup
//	wg.Add(3)
//
//	go func(j int) {
//		defer wg.Done()
//		for i := 0; i < 100; i++ {
//			<-ch1
//			fmt.Printf("%d ", j)
//			ch2 <- 2
//		}
//	}(1)
//	go func(j int) {
//		defer wg.Done()
//		for i := 0; i < 100; i++ {
//			<-ch2
//			fmt.Printf("%d ", j)
//			ch3 <- 3
//		}
//	}(2)
//	go func(j int) {
//		defer wg.Done()
//		for i := 0; i < 100; i++ {
//			<-ch3
//			fmt.Printf("%d ", j)
//			ch1 <- 3
//		}
//	}(3)
//
//	wg.Wait()
//}
//
//func test() {
//	str := "Hello"
//	fmt.Printf("%s World", str)
//}
