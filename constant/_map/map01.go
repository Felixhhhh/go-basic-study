package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m2 := map[int]struct{ x int }{
		1: {19},
		2: {x: 18}, // 可省略key、value类型标签
		3: {17},
	}
	m2[4] = struct{ x int }{x: 16}
	fmt.Println(m, m2)

	m3 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m3["a"] = 10
	m3["c"] = 30
	if v, ok := m["d"]; ok {
		println(v)
	}
	delete(m, "d")
	fmt.Println(m)

	m4 := make(map[string]int)
	for i := 0; i < 8; i++ {
		m4[string('a'+i)] = i

	}

	for i := 0; i < 4; i++ {
		for k, v := range m4 {
			print(k, ":", v, "    ")
		}
		println()
	}

	m5 := map[int]user{
		1: {"Tom", 19},
	}

	u := m5[1]
	u.age++
	m5[1] = u
	m6 := map[int]*user{
		1: {"Tom", 19},
	}

	m6[1].age++

	println(m6[1].age)
}
