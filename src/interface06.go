package main

import (
	"fmt"
	"reflect"
)

func main() {

  var varBool *bool
  fmt.Println(reflect.ValueOf(varBool).Kind())
  fmt.Println(reflect.ValueOf(varBool).Type())

  var varFloat float32
  fmt.Println(reflect.ValueOf(varFloat).Kind())
  fmt.Println(reflect.ValueOf(varFloat).Type())

  var varMap map[string]int
  fmt.Println(reflect.ValueOf(varMap).Kind()) //
  fmt.Println(reflect.ValueOf(varMap).Type()) //

  varStruct := struct{ Value int }{}
  fmt.Println(reflect.ValueOf(varStruct).Kind())
	fmt.Println(reflect.ValueOf(varStruct).Type())

}
