package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	omContact := contactInfo{email: "TEST@email.com", zipCode: 1234}
	firstUser := person{"Om", "Prakash", omContact}
	secondUser := person{
		firstName: "Ram",
		lastName:  "Bandhu",
		contactInfo: contactInfo{
			email:   "test2@email.com",
			zipCode: 12345,
		},
	}
	var thirdUser person
	var fourthUser person
	fourthUser.firstName = "Tom"
	fourthUser.lastName = "Jerry"
	fourthUser.contactInfo = contactInfo{email: "test4@email.com", zipCode: 123}
	fmt.Println(firstUser)
	fmt.Println(secondUser)
	fmt.Println(thirdUser)
	fmt.Printf("%+v", fourthUser)
	fmt.Println()

	secondUser.print()
	fmt.Println(fourthUser)
	fmt.Println()

	secondUserPointer := &secondUser
	fmt.Println(secondUserPointer)
	secondUserPointer.updateName("Ravan")
	fmt.Println()
	fmt.Println(secondUser)

	fmt.Println()
	fourthUser.updateName("Burry")
	fmt.Println(fourthUser)

	// Maps

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	fmt.Println(colors["red"])
	colors["white"] = "#000000"
	fmt.Println(colors["white"])

	for key, value := range colors {
		fmt.Println("value is", value, " and key is", key)
	}
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
