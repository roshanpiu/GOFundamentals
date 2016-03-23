//its idiomatic to return an error as the last return from functions and methods
//func testConn(target string) (rspTime float64 err error)
//if the function is success err is nil its something else if there is an error (nil is success)

package main 

import (
    "fmt"
    "os"
)

func main() {
    result, err := os.Open("text.txt")
    
    if err != nil {
        fmt.Println("Error returned was ",err)
    }else{
        fmt.Println(result)
    }
}