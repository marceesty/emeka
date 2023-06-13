package main

import "fmt"

func calculator(numb float32, sign string, numb2 float32)float32{
	var answer float32
	if sign == "+" {
		answer = numb + numb2
	}else if sign == "-"{
		answer = numb - numb2
	}else if sign == "/"{
		answer = numb / numb2
	}else if sign == "*"{
		answer = numb * numb2
	}else {
		fmt.Println("please input a valid operator")
	}
	return answer
}

func main13() {
	result := calculator(50, "/", 20)
	fmt.Println(result)
}