package main

import "fmt"

func main () {
	//dict
	webs := make(map[string]int)

	webs["first_key"] = 9
	webs["second_key"] = 90
	fmt.Println(webs["first_key"])
}
