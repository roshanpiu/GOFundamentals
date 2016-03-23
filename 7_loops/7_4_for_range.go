package main 

import (
    "fmt"
)

func main()  {
    
    //this is a slice of 3 strings
    coursesInProg := []string{"Docker", "Vmware", "Vagrant"}
    coursesCompleted := []string{"Go", "Java", "node", "Docker"}
    
    //we ignore the index value by assigning it to the underscore
    for _, i:= range coursesInProg{
        fmt.Println(i)
    }
    
    for index, i:= range coursesInProg{
        fmt.Println("index: ",index," || value: ",i)
        for _, course := range coursesCompleted {
            if i==course{
                fmt.Println("We have a clash here")
            }
            fmt.Println(course)
        }
    }
}