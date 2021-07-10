package main

import "fmt"

func main () {
	defer two() // сначала выполняется main, потом только defer
	one()
}

func one () {
	fmt.Println("1")
}

func two () {
	fmt.Println("2")
}