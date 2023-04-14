package main

//import (
//	"context"
//	"fmt"
//)
//
//func gen(ctx context.Context) <-chan int {
//	dst := make(chan int)
//	n := 1
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				fmt.Println("结束 select ")
//				return // return结束该goroutine，防止泄露
//			case dst <- n:
//				n++
//			}
//		}
//	}()
//	return dst
//}
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer fmt.Println("结束 main ")
//
//	defer cancel() // 当我们取完需要的整数后调用cancel
//
//	for n := range gen(ctx) {
//		fmt.Println(n)
//		if n == 5 {
//			break
//		}
//	}
//}
