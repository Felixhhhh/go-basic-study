package main
//
//import "fmt"
//
//func counter(out chan<- int) {
//	for i := 0; i < 100; i++ {
//		out <- i
//	}
//	close(out)
//}
//
//func squarer(out chan<- int, in <-chan int) {
//	for i := range in {
//		out <- i * i
//	}
//	close(out)
//}
//func printer(in <-chan int) {
//	for i := range in {
//		fmt.Println(i)
//	}
//}
//
//func main() {
//
//	done := make(chan int)
//	c := make(chan int)
//
//	go func() {
//		defer close(done)
//		for x := range c {
//			println(x)
//		}
//	}()
//
//	c <- 1
//	c <- 2
//	c <- 3
//	c <- 4
//	close(c)
//
//	<-done
//
//}
