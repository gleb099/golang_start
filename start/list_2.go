package main

import "fmt"

func main() {
	contactList := [3]string{"Charlie", "Bob", "Biba"} // если убрать длину, тогда можно спокойно добавлять значения

	fmt.Println(contactList)
	for i, value := range contactList {
		fmt.Printf("%d. %s\n", i+1, value)
	}

	for i := 0; i < len(contactList); i++ {
		fmt.Println(contactList[i])
	}

	//Срезы (слыйсы)
	contactList2 := []string{"Charlie", "Bob", "Biba"}
	contactList2 = append(contactList2, "Boba", "Pupa", "Lupa")
	fmt.Println(contactList2)
}
