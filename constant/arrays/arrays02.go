package main

func main() {

	//x, y := 10, 20
	//a := [...]*int{&x, &y}
	//p := &a
	//fmt.Printf("%T,%v\n", a, a)
	//fmt.Printf("%T,%v\n", p, p)

	a := [2]int{1, 3}
	p := &a
	p[1] += 10
	println(p[1])
	println(a[1])
}
