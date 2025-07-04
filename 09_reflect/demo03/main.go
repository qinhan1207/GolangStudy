package main

import (
	"fmt"
	"reflect"
)

func getTypeAndValue(arg interface{}) {
	fmt.Printf("reflect.TypeOf(arg): %v\n", reflect.TypeOf(arg))
	fmt.Printf("reflect.ValueOf(arg): %v\n", reflect.ValueOf(arg))
}
func main() {
	a := "qinhan"
	getTypeAndValue(a)
}
