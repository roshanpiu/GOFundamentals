package main 
import (
    "fmt"
    "time"
)

func main() {
    
    for timer := 10; timer >= 0; timer-- {
        
        if timer == 0 {
            fmt.Println("Boooooooooom!")
            break
        }
        fmt.Println(timer)
        time.Sleep(1 * time.Second)
    }
    
    for timer := 10; timer >= 0; timer-- {
        
        if(timer % 2 == 0){
            continue
        }
        
        if timer == 0 {
            fmt.Println("Boooooooooom!")
            break
        }
        fmt.Println(timer)
        time.Sleep(1 * time.Second)
    }
}