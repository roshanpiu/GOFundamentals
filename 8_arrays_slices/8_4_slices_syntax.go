package main 

import (
    "fmt"
)

func main() {
    //declares a slice
    //make(<type>, <length>, <capacity>)
    //<type> is the data type of the slice
    //<length> how many values will we put in it
    //<capacity> max size of the slice (max size of array that backs the slice which holds the values of the slice)
    
    //the hiden array of 10 elements are empty when this slice is initialized
    
    myCourses_1 := make([]string, 5, 10)
    
    //when we declare a slice without a capacity the capacity is same as the length of the slice
    myCourses_2 := make([]string, 5)
    
    //declaring a slice literal
    //declares a slice with type string and length 3 and capacity 3
    coursesInProg := []string{"Docker", "Vmware", "Vagrant"}
    
    
    fmt.Println("Lenght of myCourses_1:", len(myCourses_1))
    fmt.Println("Capacity of myCourses_1:", cap(myCourses_1))
    
    fmt.Println("\n")
    
    fmt.Println("Lenght of myCourses_2:", len(myCourses_2))
    fmt.Println("Capacity of myCourses_2:", cap(myCourses_2))
    
    fmt.Println("\n")
    
    fmt.Println("Lenght of coursesInProg:", len(coursesInProg))
    fmt.Println("Capacity of coursesInProg:", cap(coursesInProg))
}