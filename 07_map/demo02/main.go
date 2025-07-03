package main

import "fmt"

func printMap(cityMap map[string]string) {
	// cityMap是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key=", key, "value=", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["Jiangxi"] = "nancang"
}

func main() {
	cityMap := make(map[string]string)
	// 添加
	cityMap["Henan"] = "zhengzhou"
	cityMap["Hebei"] = "Shijiazhuang"
	cityMap["Shandong"] = "jinan"

	// 遍历
	printMap(cityMap)
	fmt.Println("-----------")
	// 删除
	delete(cityMap, "Shandong")
	// 修改
	cityMap["Hebei"] = "tangshan"

	printMap(cityMap)
	fmt.Println("-----------")
	changeValue(cityMap)
	printMap(cityMap)
}
