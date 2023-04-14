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
////func recv(c chan int) {
////	ret := <-c
////	fmt.Println("接收成功", ret)
////}
//func main() {
//	//ch := make(chan int)
//	//go recv(ch) // 启用goroutine从通道接收值
//	//ch <- 10
//	//fmt.Println("发送成功")
//
//	//ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
//	//ch <- 10
//	//fmt.Println("发送成功")
//
//	//c := make(chan int)
//	//go func() {
//	//	for i := 0; i < 5; i++ {
//	//		c <- i
//	//	}
//	//	close(c)
//	//}()
//	//for {
//	//	if data, ok := <-c; ok {
//	//		fmt.Println(data)
//	//	} else {
//	//		break
//	//	}
//	//}
//	//fmt.Println("main结束")
//
//	//ch1 := make(chan int)
//	//ch2 := make(chan int)
//	//// 开启goroutine将0~100的数发送到ch1中
//	//go func() {
//	//	for i := 0; i < 100; i++ {
//	//		ch1 <- i
//	//	}
//	//	close(ch1)
//	//}()
//	//// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
//	//go func() {
//	//	for {
//	//		i, ok := <-ch1 // 通道关闭后再取值ok=false
//	//		if !ok {
//	//			break
//	//		}
//	//		ch2 <- i * i
//	//	}
//	//	close(ch2)
//	//}()
//	//// 在主goroutine中从ch2中接收值打印
//	//for i := range ch2 { // 通道关闭后会退出for range循环
//	//	fmt.Println(i)
//	//}
//
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go counter(ch1)
//	go squarer(ch2, ch1)
//	printer(ch2)
//
//}
