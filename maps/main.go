package main

import (
	"fmt"
	"maps"
)

func main() {

	// this is first way to create a map
	// nums := make(map[string]int)
	// nums["name"] = 1
	// fmt.Println(nums)
	// fmt.Printf("a\t%v\n", nums)

	// for checking is value avail or not , if yes whats the value , you can name it anything remember this
	// m := map[string]int{"phone": 2, "mob": 935948}
	// value, ok := m["phone"]
	// if !ok {
	// 	fmt.Println("all not ok")
	// }
	// fmt.Println(m)

	m := map[string]int{"name": 2}
	m1 := map[string]int{"name": 5}

	fmt.Println(maps.Equal(m, m1))

}
