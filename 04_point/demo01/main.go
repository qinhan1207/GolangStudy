package main

import "fmt"

func swap(pa *int, pb *int) {
	// *表示对指针变量取值
	temp := *pa
	*pa = *pb
	*pb = temp
}
func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	// 定义一个指针类型的变量，并指向a的地址
	// &表示对普通变量取地址
	var p *int = &a
	fmt.Println(&a)
	fmt.Println(p)

}
