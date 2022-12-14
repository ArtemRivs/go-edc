package main

import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

func main() {

	// Реализуйте функцию generateSelfStory(name string, age int, money float64) string,
	// которая генерирует предложение, подставляя переданные данные в возвращаемую строку.
	// Например:
	// generateSelfStory("Vlad", 25, 10.00000025) // "Hello! My name is Vlad. I'm 25 y.o. And I also have $10.00 in my wallet right now."
	// Шаблон возвращаемой строки: Hello! My name is [name].
	// I'm [age] y.o. And I also have $[money with precision 2] in my wallet right now.

	// p := Person{Name: "Vasiliy", Age: 40}
	// fmt.Println("simple:", p)
	// fmt.Printf("detailed: %+v\n", p)
	// fmt.Printf("golang: %#v\n", p)

	name := "Joyce"
	age := 33
	// balance float64 = 10.44567
	var balance float64 = 10.44567

	str := fmt.Sprintf("My name is %s. I'm %d y.o. I have %.2f money", name, age, balance)
	fmt.Println(str)

}
