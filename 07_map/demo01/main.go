package main

import "fmt"

func main() {
	// 声明myMap1是一个map类型，key是string类型，value是string类型
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1是一个空map")
	}
	// 在使用map前需要使用make给map分配空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "C++"
	myMap1["three"] = "python"
	myMap1["four"] = "golang"

	fmt.Println(myMap1)
	// 第二种方式
	myMap2 := make(map[int]string)
	myMap2[1] = "shanghai"
	myMap2[2] = "beijing"
	myMap2[3] = "tianjin"
	fmt.Println(myMap2)
	// 第三种方式
	myMap3 := map[int]string{
		1: "java",
		2: "php",
		3: "python",
	}
	fmt.Println(myMap3)

}
