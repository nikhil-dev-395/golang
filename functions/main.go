// this is function

package main

import "fmt"

/*
here you can see that we are passing function as a parameter to other function
*/
func addTwoNum(a, b int) (res int) {
	res = a + b
	return res
}

func main() {
	fmt.Println(addTwoNum(2, 4))
}
