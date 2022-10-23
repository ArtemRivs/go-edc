package main

import "fmt"

var messsage string
messsage = "Usual way to define variable"
messsage2 := "Short way to define variable"

func main(){
	fmt.Print(messsage);
	fmt.Print(messsage2);
	
	firstName := "John"
	lastName := "Smith"
	
	fmt.Print(firstName + " " + lastName)
}
