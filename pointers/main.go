package main

import "fmt"

func changeNum(num *int) {
	*num = 3
	fmt.Println("at changeNum", *num)
}
func main() {
	num := 1
	changeNum(&num)
	fmt.Println("at main ", num)

}
