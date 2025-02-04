/*
* slices
 */

package main

import (
	"fmt"
)

func main() {
	/*
		* how to make slices ?
		-->  declaration_keywords variable_name [] datatype ,
		e.g. var name []string
	*/
	/*
		      ** explanation:=

		      there are two ways to make slices ,
		       1.  declaration_keywords variable_name [] datatype ,
				e.g. var name []string

		       2.  variable_name = [] datatype {}, -- here our len and cap going to dynamic
				e.g. names := []int{}

		     how to add or append a elements
		     append(slice_name , 2) or for position we need to use slice_name[index] = value


			   nums := make([]int, 2, 5)
			   	// nums := []int{}
			   	nums[0] = 3
			   	nums = append(nums, 1)
			   	nums = append(nums, 2)

			   	fmt.Println(nums)
			   	fmt.Println("length of nums", len(nums), "capacity", cap(nums))
	*/

	// how to copy slices ?

	// nums := make([]int, 0, 5)
	// nums = append(nums, 3)
	// nums2 := make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println("nums len", len(nums), "cap", cap(nums))
	// fmt.Println("nums2 len", len(nums2), "cap", cap(nums2))
	// fmt.Println("nums", nums)
	// fmt.Println("nums2", nums2)

	/*
		how to print elements between range here
	*/

	// nums := []int{1,2,3,4}
	// fmt.Println(nums[1:3])

	// how to compare slices
	// nums := []int{1, 2}
	// nums2 := []int{1, 2}
	// fmt.Println(slices.Equal(nums, nums2))

	// how to make 2d slice

	nums := [][]int{{4, 5, 6}, {1, 2, 3}}
	fmt.Println((nums))

}
