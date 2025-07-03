package main

import "fmt"

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book *Book) {
	// 指针传递
	book.title = "java"
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhangsan"

	changeBook(&book1)

	fmt.Printf("%v\n", book1)
}
