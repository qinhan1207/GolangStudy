package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("read a book")
}
func (b *Book) WriteBook() {
	fmt.Println("write a book")
}

func main() {
	// b:pair<type:Book,value:&Book{}地址>
	b := &Book{}

	// r:pair<type:,value:>
	var r Reader
	// r:pair<type:Book,value:&Book{}地址>
	r = b

	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()

}
