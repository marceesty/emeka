package main

import "fmt"

// addition adds and two numbers and returns their sum
func addition(num1, num2 int) int{
	answer := num1 + num2
	return answer
}

// subtraction minus and two numbers and returns their difference
func subtraction(num1, num2 int){
	answer := num1 - num2
	fmt.Println(answer)
}

// additionAndSubtraction adds and subtracts two numbers and returns their sum, and differences
func additionAndSubtraction(num1, num2 int)(int, int){
	answer := num1 + num2
	answer2 := num1 - num2
	return answer, answer2
}



func main12() {
	anything := addition(23, 45)
	fmt.Println(anything)

	subtraction(54, 24)

	showAdd,showSubtraction := additionAndSubtraction(76, 38)
	fmt.Println(showAdd, showSubtraction)
}