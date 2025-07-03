package main

import "fmt"

// const 来定义枚举类型
const (
	// 可以在const()添加一个关键字iota，每行iota都会累加1，第一行的iota默认值为0
	a = 10 * iota
	b
	c
)

const (
	d = 1
	e = 2
	f = 3
)

func main() {
	// 常量(只读属性)
	const length int = 10
	fmt.Printf("length: %v\n", length)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

}
