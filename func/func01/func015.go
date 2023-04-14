package main

func hello() {
	println("hello")
}

func exec(f func()) {
	f()
}

func main() {
	f := hello
	exec(f)
}
