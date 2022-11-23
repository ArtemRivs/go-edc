package main

import "fmt"

func main() {

	s := "abc"
	var step int = -511
	var str string

	for p := 0; p < len(s); p++ {

		b := int(s[p])
		b = b + step

		// fmt.Println("p:", p, "s[p]", s[p], "b", b)
		// fmt.Println("100%5:", 100%5, "100%3", 100%3, "-511%256", -511%256)

		if b > 256 || b < 0 {
			b = b % 256
		}
		str = str + string(byte(b))
	}

	fmt.Println("str:", str)

}
