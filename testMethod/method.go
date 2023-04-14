
package main

import "fmt"

type Test struct {
num int
}

func (t *Test) TestMethod(x, y int) int {
t.num = 2
fmt.Println("TestMethod修改为2：", t.num)
return x + y
}

func main() {
t := Test{}
t.num = 1
fmt.Println("传递TestMethod方法前：", t.num)

testMethod := (&t).TestMethod(1, 2)
fmt.Println("调用完方法后：", t.num)
fmt.Println(&t)

// 也可以简写成以下模式，因为编译器在底层做了优化，在编译时会自动加上(&)
testMethod2 := t.TestMethod(1, 2)
fmt.Println(testMethod)
fmt.Println(testMethod2)
}
