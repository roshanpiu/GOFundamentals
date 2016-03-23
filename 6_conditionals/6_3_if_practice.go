package main 

import "fmt"

func main() {
    firstRank := "39"
    secondRank := "619"
    
    if firstRank < secondRank {
        fmt.Println("First course if doing better than the second course")
    }else if firstRank > secondRank {
        fmt.Println("Second course if doing better than the first course")
    }else {
        fmt.Println("both are at the same rank")
    }
}