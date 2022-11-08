package main
import "fmt"
 
type mile int
type kilometer int

type person struct{
    name string
    age int
}
 

func distanceToEnemy (distance mile){
     
    fmt.Println("расстояние для противника:")
    fmt.Println(distance, "миль")
}
 
func main() {
     
    var distance mile = 5
    distanceToEnemy(distance)
    // var distance2 kilometer = 5
    // distanceToEnemy(distance2)   // ! ошибка
 
 
     // var tom person
    // tom.name = "Tom"
    // tom.age = 23

    // var tom person = person{"Tom", 23}
    
    tom := person{"Tom", 23}
    alice := person{"Alice", 30}
    

    fmt.Println("tom:", tom)
    fmt.Println("alice:", alice)
 
 
 
}
