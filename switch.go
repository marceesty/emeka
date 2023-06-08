package main

import "fmt"


func main10() {
	// a review on if, else if and else
	sex := "male"
	if sex == "male" {
		fmt.Println("His name should be called Emeka")
	}else if sex == "female"{
		fmt.Println("Her name should be called Nkiru")
	}else{
		fmt.Println("Your name should be called Confusion")
	}

	// creating a switch statement
	sex = "transgender"

	switch sex {
	case "male":
		fmt.Println("It is a boy")
	case "female":
		fmt.Println("it is a girl")
	default:
		fmt.Println("ooops, i don't know the sex")
	}
}