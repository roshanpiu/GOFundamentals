//if we define a label and not use it the compiler will throw an error

package main 


import (
    "fmt"
)

func main()  {
    
    for i:=0 ; i<10 ; i++ {
        fmt.Println(i)
        //breakPoint
        for j:=0 ; j<10 ; j++ {
            fmt.Println(j)
            for k:=0 ; k<10 ; k++ {
                fmt.Println(k)
                if(k == 2){
                    //breaks to the breakPoint label
                    //break breakPoint
                    break
                }
                            
            }
        }

    }

    
}