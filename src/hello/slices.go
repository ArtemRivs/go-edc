package main

import "fmt"

func main() {

	s := []int{10, 20, 30}

	fmt.Println("s: ", s)
	fmt.Println("s[1]: ", s[0])
	fmt.Println("s[2]: ", s[1])
	fmt.Println("s[3]: ", s[2])
	s[2] = 40

	s = append(s, 50)
	fmt.Println("s append 3: ", s)

	s[3] = 60
	fmt.Println("s change 3: ", s)

}
