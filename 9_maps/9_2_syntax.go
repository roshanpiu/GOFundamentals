package main 

import "fmt"

func main() {
    
    //declaring a map
    //the keytype must be a comparable type (==, !=)
    //all keys must be unique
    
    //2 major ways of declaring maps
    
    //declaring with make
    //here string is the keytype int is the value type
    leagueTitles := make(map[string]int)
    
    leagueTitles["Sri lanka"] = 1
    leagueTitles["Pakistan"] = 1
    leagueTitles["india"] = 1
    
    //declaring maps with map literals
    worldCups := map[string]int{
        "Sri lanka": 1,
        "Pakistan": 1,
        "India": 2,
    }
    
    fmt.Println(leagueTitles)
    fmt.Println(worldCups)

}