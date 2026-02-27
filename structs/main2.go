package main

import "fmt"

// type person struct

func something() {

	jihad := person{firstName: "jihad", age: 25, lastName: "fastlink"}

	jihad.age = 2333

	jihad.contactInfo = contactInfo{
		email:   "jihad@gmail.com",
		zipCode: 44001,
	}

	fmt.Println(jihad.contactInfo.email)
	fmt.Println(jihad.email)
}
