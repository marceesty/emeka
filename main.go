package main

import (
	"fmt"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welcome to my home page built with Golang</h1> <a href=/contact>Go to contact</a>")
	
}

func contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>This is the contact page</h1>  <a href=/>Go home</a>")
	
}

func main() {
	fmt.Println("starting our server...")
	fmt.Println("Application running on port 8080")
	
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}










// -check assignment
// -chatgpt
// -hello world
// -Go data types (int, boolean, int, float, complex, string, slice, array:has size, (map))
// -Go operators(-, +, ==, >, <, =>, =<, *, /, ++, --) (assignment, arithmetic, comparison, logical, bitwise)
// -logical (||, &&, !)
// -assignment(=, +=, -=, *=, /=, %=, &=, |=, ^=, >>=, <<=)
// -comparison operator(==, !-, >, <, >=, <=)
// -Go conditions(if, else, else if)
// -Go switch
// -Go Loops (for,if,range,)
// -Go struct
// -Go map
// -function
// -pointer