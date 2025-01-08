package main

import (
	"fmt"
	"sync"
)

func task (wg *sync.WaitGroup,ch <-chan int, num int){  // ch <-chan int is a read-only channel , Is ch chan int and ch <-chan int both are same ? No, ch chan int is a bidirectional channel and ch <-chan int is a read-only channel
	defer wg.Done()
	for {
		val , ok := <-ch  // ok is a boolean value that indicates whether the channel is open or closed
		if !ok {
			fmt.Printf("Channel %d is Closed",num)
			return
		}
		fmt.Printf("Task %d , channel %d\n",num,val)
	}

}


func main(){

	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(4)

	go task(&wg,ch,1)
	go task(&wg,ch,2)
	go task(&wg,ch,3)
	go task(&wg,ch,4)


	for i := 0; i <=100; i++ {
		ch <- i
	
	}
	close(ch) // close the channel

	wg.Wait()




}