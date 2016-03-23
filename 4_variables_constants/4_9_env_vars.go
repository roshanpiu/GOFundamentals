package main

import  "fmt"
import  "os"

func main() {
    //fmt.Println(os.Environ())
    
    // for _, env := range os.Environ() {
    //     fmt.Println(env)
    // }
    
    var name string = os.Getenv("USERNAME") 
    name_infered :=  os.Getenv("USERNAME") 
    fmt.Println("The user name is", name)
    fmt.Println("The user name is", name_infered)
}