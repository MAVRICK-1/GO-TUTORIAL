package main

import "fmt"

// Interfaces are a way to define a set of methods that a type must have in order to implement that interface.
// Interfaces are a way to define a contract for a type.

type paymentGateway interface {
	pay(payment float64)
	getName() string
}

type payment struct {
	Gateway paymentGateway
}

func (p payment) pay(payment float64) { // added a method pay to the payment struct
	fmt.Printf("Payment made using %s: %.2f\n", p.Gateway.getName(), payment)
}

type stripe struct{}

func (s stripe) pay(payment float64) { // added a method pay to the stripe struct
	fmt.Printf("Payment made using %s: %.2f\n", s.getName(), payment)
}

func (s stripe) getName() string {
	return "stripe"
}

type paypal struct{}

func (p paypal) pay(payment float64) { // added a method pay to the paypal struct
	fmt.Printf("Payment made using %s: %.2f\n", p.getName(), payment)
}

func (p paypal) getName() string {
	return "paypal"
}

func main() {
	stripeGateway := stripe{}
	paypalGateway := paypal{}

	paymentWithStripe := payment{
		Gateway: stripeGateway,
	}
	paymentWithStripe.Gateway.pay(100.0) // Payment made using stripe: 100.00

	paymentWithPaypal := payment{
		Gateway: paypalGateway,
	}
	paymentWithPaypal.Gateway.pay(200.0) // Payment made using paypal: 200.00
}
