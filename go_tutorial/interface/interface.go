package main


import "fmt"


// Interfaces are a way to define a set of methods that a type must have in order to implement that interface.
// Interfaces are a way to define a contract for a type.

type paymentGateway interface {
	pay(payment float64)
}

type payment struct {
	Gateway paymentGateway
}

func (p payment) pay(payment float64) {  // added a method pay to the payment struct 
	fmt.Println("Payment made: ", payment)
}


type stripe struct {}

func (s stripe) pay(payment float64) { // added a method pay to the stripe struct
	fmt.Println("Payment made using stripe: ", payment)
}

type paypal struct {}

func (p paypal) pay(payment float64) { // added a method pay to the paypal struct
	fmt.Println("Payment made using paypal: ", payment)
}




func main() {
	p := payment{}
	pay := payment{
		Gateway: p,
	}
	pay.Gateway.pay(100.0) // Payment made:  100

}