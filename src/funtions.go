package main

import "strconv"

func IntToString(value int) string {
	return strconv.Itoa(value)
}


func Generate(seed int) func() {
    return func() {
        fmt.Println(seed) // замыкание получает внешнюю переменную seed
        seed += 2 // переменная модифицируется
    }
    
}

func main() {
    iterator := Generate(0)
    iterator()
    iterator()
    iterator()
    iterator()
    iterator()
}
