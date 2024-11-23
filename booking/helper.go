package main

import (
	"fmt"
)

func information()(string,string,string,uint){
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets: ")
	fmt.Scan(&userTicket)


	return firstName,lastName,email,userTicket
}