package main

import "fmt"

func main() {
	var nums = make([]int, 3, 4)
	fmt.Printf("len=%d,cap=%d,slice =%v\n", len(nums), cap(nums), nums)

	// 向nums切片追加一个元素1，nums len = 4,[0,0,0,1],cap = 5
	nums = append(nums, 1)
	fmt.Printf("len=%d,cap=%d,slice =%v\n", len(nums), cap(nums), nums)

	// 向一个容量cap已经满了的slice追加元素,
	nums = append(nums, 2)
	fmt.Printf("len=%d,cap=%d,slice =%v\n", len(nums), cap(nums), nums)
}
