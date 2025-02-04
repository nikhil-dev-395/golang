package main

import "fmt"

func main() {
	/*
		we are going to see here how the `range` works in go
	*/
	/* fmt.Println("hii this is range")
	 nums := []int{6, 7, 8}
	// here we are just iterating only
	 for i := 0; i < len(nums); i++ {
	 	fmt.Println(nums[i])
	 }

	 range on array
	 for index, num := range nums {
	 	fmt.Println(index, num)
	 }

	 fmt.Println(nums)

	*/

	//  range on map = object(js)
	/*
		names := map[string]string{"name1": "nikhil", "name2": "arjun"}

		for index, val := range names {
			fmt.Println(index, val)
		}
	*/

	// range on string

	for index, val := range "nikhil" {
		fmt.Println(index, string(val))
	}
}
