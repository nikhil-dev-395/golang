/*
variadic functions are like , fmt.Println( UNLIMITED PARAMETERS ) -> it takes unlimited parameters
now lets make custom one , but in custom one we have the limited arg or par
*/

package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}
func main() {
	// fmt.Println(sum(5, 6, 3, 7, 8))

	// lets try using slice
	nums := []int{5, 6, 3, 7, 8}
	fmt.Println(sum(nums...))
}
