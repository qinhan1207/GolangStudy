package main

import "fmt"

// 返回多个返回值，匿名的
func foo1(a string, b int) (int, int) {
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	return 66, 77
}

// 返回多个返回值，有形参名称的
// r1,r2属于foo2的形参，默认为0
func foo2(a string, b int) (r1 int, r2 int) {
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

// 两个返回值是同类型的
func foo3(a string, b int) (r1, r2 int) {
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	// 给有名称的返回值变量赋值
	r1 = 10000
	r2 = 20000
	return

}

func main() {

	rt1, rt2 := foo1("qinhan", 12)
	fmt.Printf("rt1: %v\n", rt1)
	fmt.Printf("rt2: %v\n", rt2)

	rt1, rt2 = foo2("hello", 13)
	fmt.Printf("rt1: %v\n", rt1)
	fmt.Printf("rt2: %v\n", rt2)

	rt1, rt2 = foo3("fanfan~", 15)
	fmt.Printf("rt1: %v\n", rt1)
	fmt.Printf("rt2: %v\n", rt2)

}
