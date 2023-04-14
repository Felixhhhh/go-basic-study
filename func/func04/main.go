package main

func testStruct() {
	type calc struct {
		mul func(x, y int) int
	}

	x := calc{mul: func(x, y int) int {
		return x * y
	}}

	println(x.mul(2, 1))
}

func testChannel() {
	c := make(chan func(int, int) int, 2)
	c <- func(x int, y int) int {
		return x * y
	}
	println((<-c)(1, 5))
}
func main() {
	testStruct()
	testChannel()

}
