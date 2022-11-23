func foo() (int, error) { 
    // в теле функции может быть ошибка
}

result, err = foo()
if err != nil {
    // handle error
} 
