package main

import "unicode/utf8"

func main() {
	//m := _map[string]int{
	//	"abc": 123,
	//}
	//
	//key := []byte("abc")
	//x, ok := m[string(key)]
	//println(x, ok)

	//var b bytes.Buffer
	//b.Grow(1000)
	//for i := 0; i < 1000; i++ {
	//	b.WriteString("我")
	//}
	//
	//println(b.String())

	//r := '我'
	//s := string(r)  // 我
	//b := byte(r)    // 17
	//s2 := string(b) //
	//r2 := rune(b)   // 17
	//fmt.Println(s, b, s2, r2)

	s := "十.五"
	println(len(s), utf8.RuneCountInString(s))

}
