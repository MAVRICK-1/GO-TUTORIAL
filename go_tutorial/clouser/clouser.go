package main

import "fmt"

func intSeq() func() int {
	count :=0

	return func() int {
		count++
		return count
	}
}

func main() {
	nextInt := intSeq() 

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	// it increments the count each time we call nextInt and returns the incremented value because the function intSeq returns a closure.
	// The variable count is declared in the enclosing function intSeq. The anonymous function function returned by intSeq closes over the variable count to form a closure.
	// closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	newInts := intSeq()
	fmt.Println(newInts()) // 1
}	