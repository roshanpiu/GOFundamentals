// by default go passes arguments by value 
// by default when we pass a argument to a function go pass a copy of the variable
// we can pass arguments by reference using & and *

package main

import (
    "fmt"
    "reflect"
)

func main() {
    
    
    name := "Roshan"
    course := "Go Fundamentals"
    module := 10.3 
    
    
    fmt.Println("The type of name is : ", reflect.TypeOf(name))
    fmt.Println("The type of course is : ", reflect.TypeOf(course)) 
    
    //crrates a pointer to the module variable
    //ptr holds the memory address of the module variable
    ptr := &module
    
    //prints the memory address of the module variable
    fmt.Println("The memory address of module is : ", &module)
    fmt.Println("The poniter ptr points to : ", ptr)
    
    //dereferencing a pointer prints the value of module
    fmt.Println("The value of *module is : ", *ptr)
    
}