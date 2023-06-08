package main

import "fmt"


func main11() {
	type student struct{
		name string
		age int
		height float32
		complexion string
		weight int
	}
	var emeka student
	emeka.name = "Emeka"
	emeka.age = 18
	emeka.height = 6.71
	emeka.complexion = "Brown"
	emeka.weight = 100

	newStudent := student{
		name: "Amaka",
		age: 21,
	}
	fmt.Println(newStudent)

	fmt.Println(emeka)
}
