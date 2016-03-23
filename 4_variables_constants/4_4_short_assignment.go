package main

import (
    "fmt"
    "reflect"
)

func main() {
    
    //declaring varibles at function level with short assignment
    //this only works with in functions and if we are initializing variables at the same time we declare them
    //we can not declare variables if we are not using them (applies to only with in functions)
    //at the package level we can have unused variables 
    
    name := "Roshan"
    course := "Go Fundamentals"
    module := 10.3 
    
    
    fmt.Println("The type of name is : ", reflect.TypeOf(name))
    fmt.Println("The type of course is : ", reflect.TypeOf(course)) 
    fmt.Println("The type of module is : ", reflect.TypeOf(module))
    
}