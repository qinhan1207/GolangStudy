package main

import "fmt"

func main() {
	// 方法1,2,3可以声明全局变量
	// 1.声明一个变量，默认值为0
	var a int
	fmt.Printf("a: %v\n", a)
	// 2.声明一个变量，并初始化值
	var b int = 10
	fmt.Printf("b: %v\n", b)
	// 3.初始化时可省去数据类型
	var c = 3.14
	fmt.Printf("c: %v\n", c)
	// 4.省去var关键字，使用 := 只能在函数体内来声明
	d := "hello"
	fmt.Printf("d: %v\n", d)

	var e, f int = 10, 20
	fmt.Println(e, f)

	var g, h = 20, 3.14
	fmt.Println(g, h)

	l, i := 20, "nihao"
	fmt.Println(l, i)

	var (
		j int    = 10
		k string = "qinhan"
	)
	fmt.Println(j, k)

}
