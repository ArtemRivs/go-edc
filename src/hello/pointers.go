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
   
    var xx int = 4
    var pp *int  = &xx                // указатель получает адрес переменной
    fmt.Println("Address:", pp)      // значение указателя - адрес переменной x
    fmt.Println("Value:", *pp)
   
    var a int = 5
    p := &a
    fmt.Println(a,p) //a=5 p=0xc0000b2008   
   
   
}
