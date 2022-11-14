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
   
     incrementCopy := func(i int) {
      i++
  }

  increment := func(i *int) {
      (*i)++
  }

  i := 42

  incrementCopy(i)
  fmt.Println(i) // 42

  increment(&i)
  fmt.Println(i) // 43
   
  a := 1
  p := &a
  b := &p
  
  *p = 3
  **b = 5
  
  a += 4 + *p + **b
  
  fmt.Printf("%d",*p)   
   
   
   
}
