package main

import "fmt"

func add(a int, b int) int { // a and b are the arguments of the function add and `int` is the return type of the function
	return a + b
}

func mul_return(a, b int) (int, int) { // mul_return takes two arguments of type int and returns two integers
	return a + b, a * b
}

func take_func(fn func(int,int)int) int { // take_func takes a function as an argument
	return fn(3, 4)
}

func return_func() func(int, int) int { // return_func returns a function
	return add
}
func main() {
	// fmt.Println(add(42, 13)) // 55
	// fmt.Println(mul_return(42, 13)) // 55 546
	fmt.Println(take_func(add)) // 7

	fn := return_func()
	fmt.Println(fn(3, 4)) // 7

}


