package main

import "fmt"


func main(){
	// nums := []int{1,2,3,4,5,6}

	// // Push element to the array
	// nums = append(nums, 7)
	// fmt.Println("After push:", nums)

	// // Pop element from the array
	// popped := nums[len(nums)-1]
	// nums = nums[:len(nums)-1]
	// fmt.Println("After pop:", nums)
	// fmt.Println("Popped element:", popped)

	nums := make([]int, 5)

	// Push element to the array
	nums = append(nums, 1)
	nums = append(nums, 3)
	nums = append(nums, 5)
	nums =  append(nums, 4)
	fmt.Println("After push:", nums)
	// Pop element from the array
	popped := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	fmt.Println("After pop:", nums)
	fmt.Println("Popped element:", popped)

	// Remove element from the array
	index := 2
	nums = append(nums[:index], nums[index+1:]...) // ... is used to unpack the slice into individual elements 
	// nums[:index] is the elements before the index and nums[index+1:] is the elements after the index and nums[index+1:]... is used to unpack the slice into individual elements
	// logic is to append the elements before the index and after the index
	fmt.Println("After remove:", nums)
}