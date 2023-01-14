package main

import (
	"errors"
	"fmt"
)

func mypanic() {
	defer func() error {
		if r := recover(); r != nil {
			e := errors.New("Panic")
			fmt.Println("Panic!: ", r)
			return e
		}
		return nil
	}()
	var slice []string
	fmt.Println(slice[0])
}

func main() {
	mypanic()

}
