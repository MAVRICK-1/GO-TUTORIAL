package main

import (
	"fmt"
	//"slices"
)

func main() {
	// var num[]int
	// fmt.Println(num) // []

	// num := make([]int, 5) // make a slice with 5 elements parameters are (type, length, capacity)
	// fmt.Println(num) // [0 0 0 0 0]
	// println(cap(num)) // 5

	// num := make([]int,2,5) // make a slice with 2 elements and capacity of 10
	// fmt.Println(num) // [0 0]
	// fmt.Println(cap(num)) // 5

	// num = append(num, 1)
	// num = append(num, 2)
	// num = append(num, 3)
	// num = append(num, 4)
	// num = append(num, 5)
	// num = append(num, 6)

	// fmt.Println(num) // [0 0 1 2 3 4 5 6]
	// fmt.Println(cap(num)) // 10 // capacity is doubled when it reaches the limit

	// var num = []int{}

	// num = append(num, 1)
	// num = append(num, 2)
	// //num[1]=2 // panic: runtime error: index out of range [1] with length 1
	// num[1]=3
	// fmt.Println(num) // [1 3]
	// fmt.Println(len(num)) // 2
	// fmt.Println(cap(num)) // 2
	// num = append(num, 4)
	// fmt.Println(num) // [1 3 4]
	// fmt.Println(len(num)) // 3
	// fmt.Println(cap(num)) // 4 // capacity is doubled when it reaches the limit

	// Copy function

	// num1 := []int{1,2,3,4}

	// num2 := make([]int, len(num1)) // make a slice with the same length of num1

	// fmt.Println(num1,num2) // [1 2 3 4] []

	// copy(num2,num1) // copy the elements of num1 to num2 parameters are (destination, source)

	// fmt.Println(num1,num2) // [1 2 3 4] [1 2 3 4]


	//slice of slice

	// num := []int{1,2,3,4}

	// fmt.Println(num[:2]) // [1 2]  parameters are (start, end) end is exclusive
	// fmt.Println(num[2:]) // [3 4]  parameters are (start, end) end is exclusive

	// Compair slices 

	// num1 := []int{1,2,3,4}
	// num2 := []int{1,2,3,4}

	// fmt.Println((slices.Equal(num1,num2))) // true , parameters are (slice1, slice2)

	var num = [][]int{{1,2,3},{4,5,6},{7,8,9}}

	fmt.Println(num) // [[1 2 3] [4 5 6] [7 8 9]]

	fmt.Println(num[0]) // [1 2 3]

}
