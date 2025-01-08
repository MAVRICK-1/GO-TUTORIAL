package main

import (
	"fmt"
	"sync"
)


func task (id int , wg *sync.WaitGroup) { // wg is a pointer to sync.WaitGroup
	defer wg.Done() // decrement the counter ,defer will be called when the function returns
	fmt.Println("Task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)
		
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Println("Done")
}