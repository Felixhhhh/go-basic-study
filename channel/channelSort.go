package main

//
//import "sync"
//
//func main() {
//	var wg sync.WaitGroup
//	ch1 := make(chan int, 1)
//	ch2 := make(chan int, 1)
//	ch3 := make(chan int, 1)
//	ch4 := make(chan int, 1)
//	ch1 <- 1
//	wg.Add(4)
//
//	go func() {
//		for i := 0; i < 100; i++ {
//			<-ch1
//			println(1)
//			ch2 <- 2
//		}
//		wg.Done()
//	}()
//	go func() {
//		for i := 0; i < 100; i++ {
//
//			<-ch2
//			println(2)
//			ch3 <- 1
//		}
//		wg.Done()
//
//	}()
//	go func() {
//		for i := 0; i < 100; i++ {
//			<-ch3
//			println(3)
//			ch4 <- 1
//		}
//		wg.Done()
//
//	}()
//	go func() {
//		for i := 0; i < 100; i++ {
//			<-ch4
//			println(4)
//			ch1 <- 1
//		}
//		wg.Done()
//	}()
//	wg.Wait()
//}
