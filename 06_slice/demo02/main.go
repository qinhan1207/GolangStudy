package main

import "fmt"

func printArray(arr []int) {
	// 引用传递
	// _表示匿名的变量
	for _, value := range arr {
		fmt.Println(value)
	}
	arr[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4, 5} //动态数组，切片slice
	fmt.Printf("myArray type if %T\n", myArray)
	printArray(myArray)
	printArray(myArray)
}
