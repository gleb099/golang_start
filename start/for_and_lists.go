package main

import "fmt"

func main () {
	var i = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	var arr[3] int
	arr[0] = 23
	arr[1] = 97
	arr[2] = 76
	fmt.Println(arr[1])

	nums := [3]float64 {4.23, 5.23, 98.1} // Если указать не все элементы, то будут нулевые конктреного типа
	for i, value := range nums {
		fmt.Println(i, value)
	}
}
