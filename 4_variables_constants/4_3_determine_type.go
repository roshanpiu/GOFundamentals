package main

import (
    "fmt"
    "reflect"
)

var (
    name, course string //initialized with a empty string
    module float64 //initialized with a 0 for floats
)

//let the go to infer types
var (
    i_name, i_course , i_module = "Roshan", "Go Fundamentals", 10.3
   
)


func main() {
    //determining the type of variables
    fmt.Println("The type of name is : ", reflect.TypeOf(name)) 
    fmt.Println("The type of module is : ", reflect.TypeOf(module))
    
    //type infering 
    fmt.Println("The type of i_name is : ", reflect.TypeOf(i_name)) 
    fmt.Println("The type of i_course is : ", reflect.TypeOf(i_course)) 
    fmt.Println("The type of i_module is : ", reflect.TypeOf(i_module))  
    
    a := 10.111
    b := 3
    
    //go does not allow to add a floats to ints directly first we have to convert
    //here a is still a float the value is not changed
    c := int(a) + b
    fmt.Println("a + b is : ", c) 
    fmt.Println("The type of c is : ", reflect.TypeOf(c)) //when we add two ints we get an int result
}