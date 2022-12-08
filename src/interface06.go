package main

import (
	"fmt"
	"reflect"
)

type MyType struct{}

func IsNil(obj interface{}) bool {
    if obj == nil {
        return true
    }

    objValue := reflect.ValueOf(obj)
    // проверяем, что тип значения ссылочный, то есть в принципе может быть равен nil
    if objValue.Kind() != reflect.Ptr {
        return false
    } 
    // проверяем, что значение равно nil 
    //  важно, что IsNil() вызывает панику, если value не является ссылочным типом. Поэтому всегда проверяйте на Kind() 
    if  objValue.IsNil() {
        return true
    }

    return false
} 

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
