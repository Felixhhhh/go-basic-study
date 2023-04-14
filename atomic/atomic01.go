package main

//
//import (
//	"fmt"
//	"sync/atomic"
//)
//
//func main() {
//
//	// 十进制
//	var a int = 10
//	fmt.Printf("%d \n", a) // 10
//	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制
//
//	// 八进制  以0开头
//	var b int = 077
//	fmt.Printf("%o \n", b) // 77
//
//	// 十六进制  以0x开头
//	var c int = 0xff
//	fmt.Printf("%x \n", c) // ff
//	fmt.Printf("%X \n", c) // FF
//
//	var x uint64
//	x = 2
//
//	xx := ^uint64(0)
//	ss := uint64(0)
//	fmt.Println(xx, ss)
//	fmt.Printf("%b \n", xx)
//	fmt.Printf("%b \n", ss)
//	atomic.AddUint64(&x, ^uint64(0))
//	fmt.Printf("%d \n", x) // FF
//
//}
