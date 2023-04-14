package main

import (
	"fmt"
	"strconv"
)

type color byte

const (
	black color = iota
	red
	blue
)

var x = 0x100

const y = 0x200

func test(c color) {
	println(c)
}

func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	println(s)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}
func main() {

	m := mkmap()
	println(m["a"])
	s := mkslice()
	println(s[0])
	a, _ := strconv.ParseInt("1100100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	println(a, b)
	println("0b" + strconv.FormatInt(a, 2))
	println("0" + strconv.FormatInt(b, 8))

	println(&x, x)
	//println(&y, y)
	test(red)
	test(3)

	const (
		x uint16 = 120
		y
		z
	)
	fmt.Printf("%T,%v", y, y)
	fmt.Printf("%T,%v", z, z)

}
