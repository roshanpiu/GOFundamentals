package main 

import (
    "fmt"
)

func main() {
    
    //switch and case types has to match
    //when a case is matched the case staments executes the block and breaks out of the case other conditions are not validated
    //each case block has an implicit break
    //there is no automatic fallthrough behaviour 
    
    //when there is a fallthrough when a case is matched if that case has a fallthrough then the case immediately after that gets executed
    //weather it is matched or not
    switch "docker" {
    case "linux":
        fmt.Println("linux case")
    case "docker":
        fmt.Println("docker case")
        fmt.Println("docker case")
        fallthrough
    case "windows":
        fmt.Println("windows case")
    case "OSX":
        fmt.Println("OSX case")
    default:
        fmt.Println("default block")
    }
    
}