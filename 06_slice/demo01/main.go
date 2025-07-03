package main

import "fmt"

func printArray(arr [4]int) {
	// 数组为值拷贝
	for index, value := range arr {
		fmt.Println("index=", index, "value=", value)
	}
}

func main() {
	// 表示一个固定长度的数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	printArray(myArray3)

	// 查看数据的数据类型
	fmt.Printf("myArray1 type of is %T\n", myArray1)
	fmt.Printf("myArray2 type of is %T\n", myArray2)
	fmt.Printf("myArray3 type of is %T\n", myArray3)

}
