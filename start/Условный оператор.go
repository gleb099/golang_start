package main

import "fmt"

func main () {
	var age = 89
	if age < 5 {
		fmt.Println("Age < 5")
	} else if age == 5 {
		fmt.Println("age = 5")
	} else if (age > 5) && (age <= 18) {
		var grade = age - 5
		fmt.Println(grade)
	} else {
		fmt.Println("age > 18")
	}
}
