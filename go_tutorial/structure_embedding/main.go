package main

import (
	"fmt"
)

type Customer struct {
	name string
	phone string
}

type Order struct {
	OrderID int
	OrderDate string
	OrderAmount float64
	status string
	Customer
}

func main(){

	order1 := Order{
		OrderID: 1,
		OrderDate: "2021-01-01",
		OrderAmount: 100.0,
		status: "Pending",
		Customer: Customer{
			name: "John Doe",
			phone: "1234567890",
		},
	}

	fmt.Println(order1)
	fmt.Println(order1.name)
	fmt.Println(order1.phone)
}