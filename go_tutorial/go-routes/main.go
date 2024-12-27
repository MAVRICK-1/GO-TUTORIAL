package main

import (
	"fmt"
	"time"
)

func task(i int){
	fmt.Println("Task", i)
}
func main() {
	for i := 0; i < 10; i++ {
		go task(i)

		go func () {  // will see it again
			fmt.Println("Task", i)
		}()  // we have added () at the end of the function to call it immediately, what it return ? 
	}
	time.Sleep(time.Second*2)

	fmt.Println("Done")
}