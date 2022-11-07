package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 0
}
func main() {
    x := 5
    fmt.Println("begin x:", x)
    fmt.Println("begin &x:", &x)
    
    zero(&x)
    fmt.Println("end:", x)
   
    var x int = 4       // определяем переменную
    var p *int          // определяем указатель 
    p = &x              // указатель получает адрес переменной
    fmt.Println(p)      // значение самого указателя - адрес переменной x   
   
   
   
}
