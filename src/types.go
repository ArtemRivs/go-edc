package main

import "fmt"
 
type library []string
 
func printBooks(lib library){
 
    for _, value := range lib{
     
        fmt.Println(value)
    }
}
 
func main() {
     
    var myLibrary library = library{"Book1", "Book2", "Book3"}
    printBooks(myLibrary)
}
