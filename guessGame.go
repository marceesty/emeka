package main

import "fmt"


func main15() {
	fmt.Println("Welcome to my game, Please Guess a number between 1-100")

	myNumber := 51

	var Guess int
	_, err := fmt.Scan(&Guess)

	if err != nil{
		fmt.Println(err)
	}

	for count := 0; count >= 10; count++{
		if Guess == myNumber{
			fmt.Println("You guessed it right!")
			break
		}else if Guess > 60 {
			fmt.Println("Your guess is Too high!")
			
		}else if Guess < 40 {
			fmt.Println("Your guess is Too low!")
			
		}else{
			fmt.Println("Your guess is out of range!")
			
		}
	}
}