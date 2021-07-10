package main

import "fmt"

func main () {
	var a = 29
	var b = 45
	a, b = summ (a, b)
	fmt.Println(a, b)
}

func summ (n1 int, n2 int) (int, int) {
	var res int
	res = n1 + n2
	return res, n1 * n2
}