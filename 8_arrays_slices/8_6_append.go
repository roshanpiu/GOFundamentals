package main 

import "fmt"

func main()  {
    mySlice := make([]int, 1, 4)
    
    
        
    fmt.Println("Lenght of mySlice:", len(mySlice))
    fmt.Println("Capacity of mySlice:", cap(mySlice))
    
    fmt.Println("\n")
    
    for i := 1; i<15; i++ {
        
        //anytime when we add a value to the slice we call the append method
        //if the capacity is full append will double the size of the underlying array and copy all the values to the new array
        
        mySlice = append(mySlice, i)
        fmt.Println(mySlice)
        fmt.Println("Capacity of mySlice:", cap(mySlice))
    }
    
}