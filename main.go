package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.gohtml", "bootstrap.gohtml")
	if err!=nil{
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("contact.gohtml", "bootstrap.gohtml")
	if err != nil {
		panic(err)
	}
	name := "Emeka"
	tmpl.Execute(w, name)
}

func main() {
	fmt.Println("Application starting on port 8080")
	fmt.Println("Application started up successfully")
	
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

