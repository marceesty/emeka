package main

import "fmt"

func main3()  {
	var age []int

	books := []string{"things fall apart", "animal farm","eze goes to school"}

	age = []int{34, 45, 65, 78}
	fmt.Printf("the type of age declared is %T and the type of book declared is %T", age, books)
	fmt.Printf("the type of age declared is %v and the type of book declared is %v", age, books)
}