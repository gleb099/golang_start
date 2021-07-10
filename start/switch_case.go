package main

import "fmt"

func main () {
	var age = 5
	switch age {
	case 5: fmt.Println("age = 5")
	case 15: fmt.Println("age = 15")
	case 20: fmt.Println("age = 20")
	default: fmt.Println("unknown")
	}
}
