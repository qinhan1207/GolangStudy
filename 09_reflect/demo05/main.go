package main

import (
	"fmt"
	"reflect"
)

/* 结构体标签 */

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagInfo, "doc:", tagDoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
