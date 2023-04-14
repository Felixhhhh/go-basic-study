package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	type data struct {
//		*int
//		string
//	}
//	x := 100
//	d := data{
//		int:    &x,
//		string: "abc",
//	}
//	fmt.Printf("%#v\n", d)
//}

func main() {
	type user struct {
		name string `昵称`
		sex  byte   "性别"
	}
	u := user{"Tom", 1}
	v := reflect.ValueOf(u)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s:%v\n", t.Field(i).Tag, v.Field(i))
	}
}
