package main

import (
	"os"
)

//创建文件
func OpenClose() {
	f, _ := os.Open("a.txt")
	println(f.Name())
	f.Close()
	
}

func main() {
	OpenClose()
}
