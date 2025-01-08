package main

import "fmt"

func add(nums ...int) int {
	sum:=0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func interface_func(val ...interface{}) { // interface{} is a type that can hold values of any type
	for _, v := range val {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println(add(1,2,3,4,5)) // 15
	fmt.Println(add(1,2,3,4,5,6,7,8,9,10)) // 55
	fmt.Println(add(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15)) // 120

	interface_func(1,"hu","hello") // 1 hu hello
	
}