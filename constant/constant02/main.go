package main

import "fmt"

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func main() {

	type (
		user struct {
			name string
			age  uint8
		}
		event func(string) bool
	)
	u := user{"abc", 16}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	var cc = f("abc")
	println(cc)

	fmt.Printf("%b\n", read)
	fmt.Printf("%b\n", exec)

}
