package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is called")
	fmt.Printf("%v\n", u)
}

func DoFieldAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is ", inputType.Name())
	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is ", inputValue)
	// 分别通过type获取里面的字段
	// 1.获取interface的reflect.Type,通过Type得到NumField,进行遍历
	// 2.得到每个field，数据类型
	// 3.通过filed有一个Interface方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
	}
	// 通过type获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)
	}
}

func main() {
	user := User{1, "qinhan", 22}
	DoFieldAndMethod(user)

}
