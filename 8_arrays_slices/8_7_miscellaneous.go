package main 

import "fmt"

func main() {
    mySlice_1 := []int{1,2,3,4,5,6}
    mySlice_2 := []int{7,8,9,10}
    
    
    //mySlice refereces the entire underlying array not just the frst value of it
    
    //its illegal in go to append slice to a slice
    //we can only append additional values to a slice with the same data type
    
    fmt.Println(mySlice_1)
    
    for _, i:= range mySlice_1{
        fmt.Println(i)
    }
    
    //the leagal way to append 2 slices with the same data type
    //here we append a collection of ints to the mySlice_1 there for it is legal
    newSlice := append(mySlice_1,mySlice_2...)
    fmt.Println(newSlice)
    
}