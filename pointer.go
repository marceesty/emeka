package main

import (
	"fmt"
)


func main14() {
	b := 50
	var a *int
	a = &b
	fmt.Println(*a)

	*a = 70

	fmt.Println(b)
}