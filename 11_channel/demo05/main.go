package main

import (
	"fmt"
)

func fn(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// 如果c可写，case就会进来
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fn(c, quit)
}
