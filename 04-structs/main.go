package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	isMarried bool
	hobbies   []string
	contactInfo
}

func main() {
	something()

	jihad := person{
		firstName: "Jihad",
		contactInfo: contactInfo{
			email:   "jihad@gmail.com",
			zipCode: 44001,
		},
	}

	jihadPointer := &jihad
	jihadPointer.updateName("hooo")
	jihadPointer.print() // Go automatically converts reference (pointer) to actual value
	println()

	// Go automatically converts jihad to pointer (reference) since we use *person receiver
	jihad.updateName("Jojo")
	jihad.print()

	name := "jihad"
	namePointer := &name
	fmt.Println(&namePointer)
	updateName(namePointer)
}

func updateName(namePointer *string) {
	fmt.Println(&namePointer)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName //Go automatically change the receiver p to (*p)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
