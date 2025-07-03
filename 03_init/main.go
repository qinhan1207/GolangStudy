package main

import (
	// _ "github.com/qinhan1207/GolangStudy/03_init/lib1" // 表示这个包可以不被使用
	// mylib2 "github.com/qinhan1207/GolangStudy/03_init/lib1" // 给包起别名
	"github.com/qinhan1207/GolangStudy/03_init/lib1"
	"github.com/qinhan1207/GolangStudy/03_init/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
