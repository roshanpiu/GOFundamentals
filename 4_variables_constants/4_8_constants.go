//constants are immutable once we assign a value we can not change it

//here we can not use := the short hand notation

package main

import (
    "fmt"
)

const speedOfLightMph = 186000

func main() {
    fmt.Println("The speed of light is ", speedOfLightMph)
}