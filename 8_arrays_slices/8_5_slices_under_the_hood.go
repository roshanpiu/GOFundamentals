package main 

import (
    "fmt"
)

func main() {
    mySlice := []int{1,2,3,4,5,6,7,8,9,10}
    
    //slices are by default references
    //slices are passed by reference by default 
    //when a slice is passed to a function it is passed as a reference
    
    //index starts from 0
    //length and capacity starts form 1
    
    fmt.Println(mySlice[4])
    
    mySlice[1] = 0
    fmt.Println(mySlice)
    
    //sliceOfSlice references values 2 to 5 from mySlice
    //when slicing a slice lefthand operator is inclusive righthand operator is exclusive
    //here sliceOfSlice gets values of index 2,3,4
    sliceOfSlice_1 := mySlice[2:5]
    
    sliceOfSlice_2 := mySlice[:5] //gets values from index 0 to 5
    sliceOfSlice_3 := mySlice[2:] //gets values from index 2 to last index
    
    fmt.Println(sliceOfSlice_1)
    fmt.Println(sliceOfSlice_2)
    fmt.Println(sliceOfSlice_3)

}