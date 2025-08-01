/* 无缓冲channel通信 */
package main

import "fmt"

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行")

		c <- 666 //将666发送给c
	}()

	num := <-c // 从c总接收数据，并赋值给num

	fmt.Println("num=", num)
	fmt.Println("main goroutine结束")
}
