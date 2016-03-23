package main

import (
    "fmt"
)

// when we declare variables at package level we have to use the var keyword
var (
    name, course string //initialized with a empty string
    module float64 //initialized with a 0 for floats
)


func main() {
    fmt.Println("The name is : ", name) 
    fmt.Println("The module is : ", module) 
}