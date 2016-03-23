package main 

import (
    "fmt"
)

func main() {
    bestFinish := bestLeagueFinishes(10, 20, 30 , 4, 30, 2, 78, -12)
    fmt.Println(bestFinish)
}

//... tells our function to accept any number of ints
//this function returns an int
func bestLeagueFinishes(finishes ...int) int {
    
    //the values passed in as arguments becomesa slice of ints
    
    best := finishes[0]
    
    for _, i := range finishes {
        if i < best {
            best = i
        }
    }
    
    return best
}