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
}
