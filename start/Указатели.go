package main

import "fmt"

//Указатели позволяют передавать в функции не само значение переменной, а именно ее адрес
func main() {
	var x = 0
	pointer(&x) // амперсант для передачи
	fmt.Println(x)
}

func pointer (x *int) { // через звездочу передаю ссылку
	*x = 2
}
