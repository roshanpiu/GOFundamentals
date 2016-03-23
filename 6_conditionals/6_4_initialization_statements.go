// we can use these to declare things that we are going to use in the if block
// the simple statements are scoped only to if else if blocks

//if one conditions becomes true in if else statements then we do not check the other conditions
//we execute the code inside the code block then exit

package main 

import "fmt"

func main() {
    
    if firstRank, secondRank := 39, 614; firstRank < secondRank {
        fmt.Println("First course if doing better than the second course")
    }else if firstRank > secondRank {
        fmt.Println("Second course if doing better than the first course")
    }else {
        fmt.Println("both are at the same rank")
    }
}