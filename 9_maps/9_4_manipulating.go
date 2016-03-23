package main 

import "fmt"

func main() {
    
    worldCups := map[string]int{
        "Sri lanka": 1,
        "Pakistan": 1,
        "India": 2,
    }
    
    //if the key exists it value gets updated otherwise a new item is added
    
    //viewing
    fmt.Println(worldCups["Sri lanka"])
    
    //updating
    worldCups["Sri lanka"] = 2
    fmt.Println(worldCups["Sri lanka"])

    //adding
    worldCups["West indies"] = 2
    fmt.Println(worldCups["West indies"])
    
    //deleting
    fmt.Println(worldCups)
    delete(worldCups, "Sri lanka")
    fmt.Println(worldCups)
    
    

}