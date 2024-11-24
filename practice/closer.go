package main

import "fmt"

func increment() func()int {

	count :=0 

	return func() int {
		count++ 
		return count
	}

}


func main(){
	nextInt := increment()

	fmt.Println(nextInt())
	fmt.Println(nextInt())


}