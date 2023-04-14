package main

//
//import "fmt"
//
//func main() {
//
//	ch1 := make(chan int)
//	ch2 := make(<-chan int)
//	ch3 := make(chan<- int)
//
//	ch2 = ch1               // 可以将通道类型的值分配给只读通道类型的变量
//	ch3 = ch1               // 可以将通道类型的值分配给只写通道类型的变量
//	fmt.Printf("%T\n", ch2) // 输出：<-chan int
//	fmt.Printf("%T\n", ch3) // 输出：chan<- int
//
//	//ch1 := make(chan int)
//	//ch2 := make(<-chan int)
//	//ch3 := make(chan<- int)
//	//
//	//ch1 = ch2 // 不能将只读通道类型的值分配给通道类型的变量
//	//ch3 = ch2 // 不能将只读通道类型的值分配给只写通道类型的变量
//	//ch1 = ch3 // 不能将只读通道类型的值分配给通道类型的变量
//	//ch2 = ch3 // 不能将只读通道类型的值分配给只写通道类型的变量
//	//fmt.Printf("%T\n", ch1)
//}
