package main

import (
	"fmt"
	"testing"
)

func test() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}
func testCap() map[int]int {
	m := make(map[int]int, 1000)
	fmt.Println(m)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}

func main() {
	resTest := testing.Benchmark(BenchmarkTest)
	fmt.Printf("BenchmarkTest \t %d, %d ns/op,%d allocs/op, %d B/op\n", resTest.N, resTest.NsPerOp(), resTest.AllocsPerOp(), resTest.AllocedBytesPerOp())
	resTest = testing.Benchmark(BenchmarkTestCap)
	fmt.Printf("BenchmarkTestCap \t %d, %d ns/op,%d allocs/op, %d B/op\n", resTest.N, resTest.NsPerOp(), resTest.AllocsPerOp(), resTest.AllocedBytesPerOp())
}
