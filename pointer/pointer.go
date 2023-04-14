package main

import (
	"fmt"
	"unsafe"
)

type Ins struct {
	x bool  // 1个字节
	y int32 // 4个字节
	z byte  // 1个字节
}

func main() {

	ins := Ins{}
	fmt.Printf("ins size: %d, align: %d\n", unsafe.Sizeof(ins), unsafe.Alignof(ins))

	a := int(1)

	b := (*int64)(unsafe.Pointer(&a)) // 将 *int 先转为 Pointer，再转为 *int64

	c := uintptr(unsafe.Pointer(&a)) // 将 *int 先转为 Pointer，再转为 uintptr

	fmt.Printf("%p\n", b) // 打印地址 0xc0003cdbb0
	fmt.Printf("%x\n", c) // 地址 c0002124b8

	type T struct {
		a string
		b int
	}
	t := T{a: "abc", b: 1}

	/*
		1. 将 t 的地址转为 Pointer：符合第一种
		2. 将 Pointer 转为 uintptr 后得到地址的整数值：符合第四种
		3. 加上 t.b 的offset，得到 t.b 的地址整数值：uintptr是整数，可以直接相加
		4. 将 uintptr 转为 Pointer：符合第三种
		5. 将 Pointer 转为 *int :符合第二种
		6. 最后解引用，得到具体的值
	*/
	d := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&t)) + unsafe.Offsetof(t.b)))
	fmt.Println(d) // 1
}
