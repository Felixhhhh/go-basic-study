package main

func main() {

	x := 100
	i := *int
	p := (&x) // 错误	Cannot convert an expression of the type '*int' to the type 'int'
	println(p)      // 		Invalid indirect of 'int(&x)' (type 'int'
}
