package main

import "fmt"

func main() {
	// 1.声明slice1是一个切片，长度为4
	slice1 := []int{1, 2, 3, 4}
	fmt.Printf("slice1:%v,len:%d\n", slice1, len(slice1))
	// 2.声明slice2是一个切片，但是并没有分配空间
	var slice2 []int
	slice2 = make([]int, 3) // 开辟三个空间，默认值为0
	fmt.Printf("slice2:%v,len:%d\n", slice2, len(slice2))
	// 3.声明slice3是一个切片，同时给slice3分配5个空间
	var slice3 []int = make([]int, 5)
	fmt.Printf("slice3:%v,len:%d\n", slice3, len(slice3))
	// 4.通过 := 来推断出slice4是一个切片
	slice4 := make([]int, 5)
	fmt.Printf("slice4:%v,len:%d\n", slice4, len(slice4))

	// 判断一个slice是否为空
	var slice5 []int
	if slice5 == nil {
		fmt.Println("slice5是一个空切片")
	} else {
		fmt.Println("slice5是有空间的")
	}
}
