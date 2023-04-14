package main

func test(f func()) {
	f()
}

func test01(x, y int) func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}

}
func main() {
	//直接执行
	func(s string) {
		println(s)
	}("hello,world")

	//赋值给变量
	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 2))

	//作为参数
	test(func() {
		println("hello,world")
	})

	//作为返回值
	con := test01
	println(con(1, 6))

}
