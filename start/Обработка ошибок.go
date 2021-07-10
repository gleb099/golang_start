package main

import (
	"errors"
	"fmt"
)

const winePrice = 100

func main() {
	change, err := buyWine(45, 19)
	if err != nil {
		fmt.Println("Purchase unsuccessful.")
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Your change - %d$, have a good day bro", change)
	}
}

func buyWine(age, moneyAmount int) (int, error) {
	if age >= 18 && moneyAmount >= winePrice {
		return moneyAmount - winePrice, nil
	} else if age < 18 {
		return moneyAmount, errors.New("You have to go to school bro:(")
	} else {
		return moneyAmount, errors.New("You dont have enought money bro:(")
	}
}
