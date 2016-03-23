package main 

import "fmt"

func main() {
    //declared using the literal form
    worldCups := map[string]int{
        "Sri lanka": 1,
        "Pakistan": 1,
        "India": 2,
    }
    
    //go uses a random starting pointing to iterate over a map
    //this is because maps are unordered lists
    for key, value := range worldCups {
        fmt.Println("Key : ",key," Value : ", value)
    }


}