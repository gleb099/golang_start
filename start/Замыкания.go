package main

import "fmt"

//Замыкания, в отличие от обычных функций, пишутся внутри других функций
func main () {
	var num = 3
	multiple := func() int {
		num *= 2
		return num
	}
	fmt.Println(multiple())
}

