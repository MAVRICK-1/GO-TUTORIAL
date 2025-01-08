package main


import (
	"fmt"
)

// func printTask(i int){
// 	fmt.Println("Task", i)
// }

func printTask [T comparable](val []T){ // T comparable is a constraint, it means that T must be a comparable type or [T int | string | float64]
	for _, v := range val {
		fmt.Println(v)
	}
}

// If for both arry and slice we can use [T any] instead of [T comparable]
// func printTask [T any](val []T){
// 	for _, v := range val {
// 		fmt.Println(v)
// 	}
// }

// print array and single value using generic
func printTaskAny[T any](val ...T){  // (val ...T) is a variadic parameter, it means that it can accept multiple values of type T
	for _, v := range val {
		fmt.Println(v)
	}
}

func main() {
	names := []string{"John", "Doe", "Jane", "Doe"}
	printTask(names)
	

	val := 3

	printTaskAny(val)


	val2 := []int{1, 2, 3, 4, 5}
	printTask(val2)


}