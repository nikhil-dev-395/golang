package main

import "fmt"

func main() {
	/*
		we are going to see here how the `range` works in go
	*/
	fmt.Println("hii this is range")
	nums := []int{6, 7, 8}
	// here we are just iterating only
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Println(nums)
}
