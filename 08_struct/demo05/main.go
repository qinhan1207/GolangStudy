package main

import "fmt"

// interfaces是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)

	// interface{}如何区分 此时引用的底层数据类型到底是什么？
	// 给interface{}提供类型断言的机制
	value, ok := arg.(Book)
	if !ok {
		fmt.Println("arg is not Book type")
	} else {
		fmt.Println("arg is Book type,value=", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"qinhan"}
	myFunc(book)
	myFunc(100)
	myFunc(3.14)
}
