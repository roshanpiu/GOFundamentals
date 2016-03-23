package main

import "fmt"

func main() {
    
    //every feild must be unique
    type courseMeta struct {
        Author string
        Level string
        Rating float64
    }
    //structs initializes with a zero value when we declare them
    //there are 3 ways to declare a struct
    
    var course_1 courseMeta //first way
    
    //this method gives us a pointer
    course_2 := new(courseMeta) //second way
    
    course_1.Author = "piumal"
    course_1.Level = "advanced"
    course_1.Rating = 5
    
    course_2.Author = "roshan"
    course_2.Level = "Beginner"
    course_2.Rating = 4.5
    
    //literal way of declaring and initializing structs
    //third way
    course_3 := courseMeta{
        Author: "Roshan Piumal",
        Level: "Intermediate",
        Rating: 5,
    }
    
    fmt.Println(course_1, course_2, course_3)
    
    course_3.Rating = 4.5
    
    
    fmt.Println("Author of course_3 : ",course_3.Author)
    fmt.Println("Rating of course_3 : ",course_3.Rating)
}