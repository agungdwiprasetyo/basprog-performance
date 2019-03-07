package main

import (
    "fmt"
)

// type Animalizer interface {
//     GetColour() string
// }

// type Animal struct {
//     Colour string
//     Name   string
// }

// type Dog struct {
//     Animal
// }

// func GetColour(a *Animal) string {
//     return a.Colour
// }

// func PrintColour(a Animalizer) {
//     fmt.Println(a.GetColour())
// }

func main() {
    sum := 0
    n := 1000000000
    fmt.Print("n= ")
    fmt.Println(n)
    for i := 0; i < n; i++ {
        sum += i
    }
    fmt.Print("sum= ")
    fmt.Println(sum)
}