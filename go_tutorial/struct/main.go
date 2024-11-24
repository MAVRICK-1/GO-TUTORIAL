package main

import (
	"fmt"
	"structs"
	"time"
)

type Order struct {
	OrderID int
	OrderDate string
	OrderAmount float64
	status string
	timeStamp time.Time
}

func  newOrder (id int, date string, amount float64, status string) (*Order) { // func (o *Order) is a method it is a function with a receiver example o *Order here o is the receiver 
	Order := Order{
		OrderID: id,
		OrderDate: date,
		OrderAmount: amount,
		status: status,
	} 
	return &Order
}

func (o *Order) ChangeStatus(status string)  { // func (o *Order) is a method it is a function with a receiver example o *Order here o is the receiver 
	o.status = status //we use pointer because we are changing the value of the struct
}

func (o *Order) 	TimeStamp(){ // func (o *Order) is a method it is a function with a receiver example o *Order here o is the receiver
	o.timeStamp = time.Now().Local()
}

func (o *Order) GetStatus() string {
	return o.status
}

func (o *Order) GetAmount() float64 {  //float64 in which 64 reffers to the number of bits
	return o.OrderAmount
}

func main(){

	// order1 := Order{
	// 	OrderID: 1,
	// 	OrderDate: "2021-01-01",
	// 	OrderAmount: 100.0,
	// 	status: "Pending",
	// }

	// fmt.Println(order1)

	// order1.ChangeStatus("Completed")
	// order1.TimeStamp()
	// fmt.Println(order1.GetStatus())
	// fmt.Println(order1.GetAmount())
	// fmt.Println(order1)

	order2 := newOrder(2, "2021-01-02", 200.0, "Pending")
	fmt.Println(order2)
	order2.ChangeStatus("Completed")
	order2.TimeStamp()
	fmt.Println(order2.GetStatus())
	fmt.Println(*order2)


	// inlining struct
	language := struct{
		Name string
		Version float64
	}{
		Name: "Go",
		Version: 1.15,
	}
	fmt.Println(language)

}