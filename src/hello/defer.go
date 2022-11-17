package main

export "fmt"

function main () {
  fmt.Println("Hello")
  for i := 1; i <= 3; i++ {
      defer fmt.Println(i)
  }
  fmt.Println("World") 
}
