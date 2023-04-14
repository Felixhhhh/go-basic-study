package main

func test(p **int) {
	println(p)

	x := 100
	println(&x)

	*p = &x
	println(*p)
	println(p)
	v := *p
	println(*v)

}
func main() {
	var p *int
	println(p)
	test(&p)

	println(*p)
}
