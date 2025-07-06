/* 有缓冲channel */
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 带有缓冲的channel

	fmt.Println("len:", len(c), "cap:", cap(c))

	go func() {
		defer fmt.Println("子go协程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行,发送的元素i=", i, "len:", len(c), "cap:", cap(c))

		}

	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c // 从c中接收并赋值num
		fmt.Println("num=", num)
	}

	fmt.Println("main 结束")
}
