package main

import  "fmt"
import  "strings"

func main() {
    module := "function basics"
    author := "Roshan Piumal"
    
    fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string)  {
    
    //this function signature can also be written like 'func converter(module string, author string) (s1 string, s2 string)'
    //s1 and s2 are return values of type string
    
    module = strings.Title(module)
    author = strings.ToUpper(author)
    
    return module, author
}