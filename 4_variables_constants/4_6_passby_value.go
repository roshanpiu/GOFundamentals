package main

import (
    "fmt"
)

func main() {
      
    name := "Roshan"
    course := "Go Fundamentals"
    
    fmt.Println("\nHi ", name, "You are currently watching ", course)
    
    changeCourse(course)
    
    fmt.Println("\nHi ", name, "You are still currently watching ", course)
}

//this function returns a string and accepts a string parameter named course
func changeCourse(course string) string  {
    
    // here we are to using :=  because we not declaring a new variable
    // here we are just assining a different value to the course variable 
    course = "Go Advanced"
    
    fmt.Println("\nChanged the course to ", course)
    
    return course
}