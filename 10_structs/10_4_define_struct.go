package main

import "fmt"

func main() {
    
    //every feild must be unique
    type courseMeta struct {
        Author string
        Level string
        Rating float64
    }
    //initializes with a zero value
    var course_1 courseMeta
    
    //this method gives us a pointer
    course_2 := new(courseMeta)
    
    course_1.Author = "piumal"
    course_1.Level = "advanced"
    course_1.Rating = 5
    
    course_2.Author = "roshan"
    course_2.Level = "Beginner"
    course_2.Rating = 4.5
    
    //literal way of declaring and initializing structs
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