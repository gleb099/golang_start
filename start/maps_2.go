package main

import "fmt"

func main() {
	contactList := map[string]string{
		"Charlie": "23476343",
		"Bob": "8965455434",
		"Biba": "313476344"}

	for name, phone := range contactList {
		fmt.Println(name, phone)
	}

	delete(contactList, "Bob")
	fmt.Println(contactList["Bob"]) // оно возвращает пустое значение
}
