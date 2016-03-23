package main 

import (
    "fmt"
)

func main() {
    
    //switch and case types has to match
    //when a case is matched the case staments executes the block and breaks out of the case other conditions are not validated
    //each case block has an implicit break
    //there is no automatic fallthrough behaviour 
    switch "docker" {
    case "linux":
        fmt.Println("linux case")
    case "docker":
        fmt.Println("docker case")
        fmt.Println("docker case")
    case "windows":
        fmt.Println(" windows case")
    default:
        fmt.Println("default block")
    }
    
}