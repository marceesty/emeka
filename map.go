package main

import (
	"fmt"
)

func main5()  {
	books := make(map[string]string)

	books["1"] = "learning golang"
	books["2"] = "learning Python"

	// fmt.Printf("the type of age declared is %T", books)
	// fmt.Printf("the type of age declared is %v", books)

	// numbers := []int{2, 4, 6, 7, 9, 8, 1,10}
	// name := "nneka"
	// age := 30

	// var cars map[string]string
	cars := map[string]string{"honda": "civic", "toyota":"venza", "benz":"AMG"}
	// cars["honda"] = "civic"
	// cars["toyota"] = "venza"

	for _, car := range cars{
		fmt.Println(car)
	}

}