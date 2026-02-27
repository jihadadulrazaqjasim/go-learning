package main

import (
	"fmt"
)

// PaymentMethod defines behavior for all payment types
type PaymentMethod interface {
	Process(amount float64) error
	GetName() string
}

type CreditCard struct {
	cardNumber string
	holderName string
}

type Paypal struct {
	email string
}

type Cash struct{}

// ------------------ GetName Methods ------------------

func (c CreditCard) GetName() string {
	return "Credit Card"
}

func (p Paypal) GetName() string {
	return "PayPal"
}

func (c Cash) GetName() string {
	return "Cash"
}

// ------------------ Process Methods ------------------

// Improvement: Consistent validation and proper money formatting
func (c Cash) Process(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid payment amount: %.2f", amount)
	}

	fmt.Printf("Processing %.2f in Cash\n", amount)
	return nil
}

func (c CreditCard) Process(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid payment amount: %.2f", amount)
	}

	// Safe slicing to avoid panic
	last4 := c.cardNumber
	if len(last4) > 4 {
		last4 = last4[len(last4)-4:]
	}

	fmt.Printf("Processing %.2f using Credit Card ending in %s\n", amount, last4)
	return nil
}

func (p Paypal) Process(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid payment amount: %.2f", amount)
	}

	fmt.Printf("Processing %.2f via PayPal account %s\n", amount, p.email)
	return nil
}

// ------------------ Checkout ------------------

// Improvement: Proper error handling and no redundant newline usage
func Checkout(pm PaymentMethod, amount float64) {
	fmt.Println("Payment Method:", pm.GetName())

	err := pm.Process(amount)
	if err != nil {
		fmt.Println("Error processing payment:", err)
		fmt.Println()
		return
	}

	fmt.Println("Payment processed successfully")
	fmt.Println()
}

func main() {
	card := CreditCard{
		cardNumber: "1234567890123456",
		holderName: "Jihad Abdulrazaq",
	}

	paypal := Paypal{
		email: "jihad@gmail.com",
	}

	cash := Cash{}

	Checkout(card, 10)
	Checkout(paypal, 200)
	Checkout(cash, 20)

	// Test error case
	Checkout(cash, -1)

	// Shape interface demo
	sq := square{sideLength: 2}
	printArea(sq)

	rect := rectangular{length: 2, width: 3}
	printArea(rect)
}
