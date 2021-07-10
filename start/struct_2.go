package main

import "fmt"

func main() {
	charlie := Contact {
		"Charlie",
		"Potter",
		"8982342345"}
	fmt.Printf("Контакт 1:\n%s\n%s\n%s\n", charlie.first_name, charlie.last_name, charlie.phone_number)
	fmt.Println(charlie.hello())
}

type Contact struct {
	first_name string
	last_name string
	phone_number string
}

func (contact *Contact) hello () string {
	return "Hello, " + contact.first_name + "!"
}