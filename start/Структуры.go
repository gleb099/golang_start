package main

import "fmt"

func main() {
	bob := Cats {
		"Bob",
		7,
		0.87}
	fmt.Println("Bob age is", bob.age)
	fmt.Println(bob.test())
}

type Cats struct {
	name string
	age int
	happiness float64
}

func (cat *Cats) test () float64 {
	return float64(cat.age) * cat.happiness
}