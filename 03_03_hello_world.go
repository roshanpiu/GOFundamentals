//'package main' this is the package declaration
//'package main' tell the compiler that this is a stand executable alone program rather than a shared library
//'package main' is the entry point to our app when coupled with the main function

package main

//go is case sensitive

//importing the format and runtime packages
//standerd libraries can be imported like this using the short name
//when importing our own packages we need to give explicit paths
import (
    "fmt"
    "runtime"
)


func main() {
    fmt.Println("Hello World!", runtime.GOOS) 
}