package main

func main() {
	x, y := 1, 2
	defer func(a int) {
		println("defer x,y =", x, y)
	}(x)

	x += 100
	y += 200
	println(x, y)

}
