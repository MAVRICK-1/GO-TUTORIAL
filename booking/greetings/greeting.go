package greetings

import (
	"fmt"
)
func Greetings(conferenceName string, conferenceTickets uint, remainingTickets uint ){
	fmt.Printf("Conference is %T, remaining Tickets are %T, conferenceTickets is %T\n", conferenceName, remainingTickets, conferenceTickets)
	//to see all the types
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Choose your tickets wisely!")
}