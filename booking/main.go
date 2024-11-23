package main

import (
	"fmt"
	"strings"
	"booking/greetings"
)


func main() {

	var conferenceName string = "Go conference"
	var conferenceTickets uint = 50
	var remainingTickets uint = 50

	

	bookings := []string{}

	greetings.Greetings(conferenceName,conferenceTickets,remainingTickets)
	for {

		firstName,lastName,email,userTicket:=information()
		var userData=make(map[string]string)
		userData["firstName"]=firstName



		if remainingTickets < userTicket {
			fmt.Printf("Cannot complete your order, only %v tickets left\n", remainingTickets)
			continue
		}

		isValidFirstName := len(firstName) > 0
		isValidLastName := len(lastName) > 0
		isValidEmail := strings.Contains(email, "@")

		if isValidFirstName && isValidLastName && isValidEmail {
			remainingTickets = remainingTickets - userTicket
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Name: %v\nEmail: %v\nRemaining Tickets: %v\n", firstName+" "+lastName, email, remainingTickets)
			fmt.Println("Current bookings:", bookings)

			firstNames := []string{}
			for _, booking := range bookings {
				var name = strings.Fields(booking)
				firstNames = append(firstNames, name[0])
			}
			fmt.Println("First names of current bookings:", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out!")
				break
			} else {
				fmt.Printf("Remaining tickets: %v\n", remainingTickets)
			}
		} else {
			fmt.Println("Invalid input! Please try again.")
		}

	}

	//switch case 
	city:="NewYork"

	switch city {
	case "NewTown","India","China":
		fmt.Println("Your Application process form B")
	case "NewYork","china":
		fmt.Println("Got it")
	case "Hooghly":
		fmt.Println("Do it")
	default:
		fmt.Println("not found")
	}

	
}
