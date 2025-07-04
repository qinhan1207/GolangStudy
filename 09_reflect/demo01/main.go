package main

import "fmt"

func main() {

	// pair<statictype:string,value:"qinhan">
	var a string = "qinhan"

	// pair<type:string,value:"qinhan">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Printf("str: %v\n", str)

}
