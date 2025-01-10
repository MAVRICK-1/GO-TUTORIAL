package main

import (
	"fmt"
	"time"
)

// <-chan (only read)  chan<- only write

//synchronizer
func synchronizer (done chan<- bool) {
	defer func (){ done <- true}()  // this will always run
	fmt.Println("Processing.....")
}


func results (ch chan int, num int, num2 int){
	ch <- (num+num2)
	
}

func emailSender(emails chan string, done chan bool) {

	defer func () {done <- true}()

	for email := range(emails){
		fmt.Printf("Sending mail to %s\n", email)
		time.Sleep(time.Second)  // to simulate the time taken to send the email
	}

}

func main(){

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func () {
		chan1 <- 10
	}()

	go func () {
		chan2 <- "ping"
	} ()


	for i:=0 ; i<2 ; i++ {
		select{
		case chan1Val := <-chan1:
			fmt.Printf("Messag from chanel1 %d\n", chan1Val)
		case chal2Val := <-chan2:
			fmt.Printf("Message from Chanel2 %s\n", chal2Val)
		}
	}




	// emails := make(chan string, 100) // buffered channel , buffer Channel is used to store the data in the channel and it will not block the sender until the buffer is full
	// done := make(chan bool)

	// go emailSender(emails,done)

	// for i:=1 ; i<=30 ; i++ {

	// 	emails <- fmt.Sprintf("%d@gmail.com",i)
	// }

	// close(emails) // close the channel

	// <-done // it will block the main function until the emailSender goroutine is completed

	// fmt.Println("email send successfully")





	// done := make(chan bool)

	// go synchronizer(done)

	// <-done // it will block it till here untill the process completes we don't have to use the wait group

	// Important

	// use Wait group for multiple process and channel for single process




	// ch := make(chan int)

	// go results(ch,9,6)

	// res := <-ch  // why blocking ? because we are trying to read from the channel before the value is written to the channel by the goroutine
	// fmt.Println(res)


	
}