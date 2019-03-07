package main

import (
    "fmt"
    "time"
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
    n := 1000000000
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = i
    }
    has :=0
    val := 999999998

    start := time.Now()
    for i := 0; i < len(arr); i++ {
        if val==arr[i]{
            has = i
            break
        }
    }
    elapsed := time.Since(start)
    fmt.Println("mainstream -> ",elapsed, "has -> ", has)
    start = time.Now()
    nm := len(arr)
    f, m, l := 0, 0, nm-1
    for f <= l {
        m = (f+l)/2
        if val > arr[m]{
            f = m +1
        } else if val < arr[m] {
            l = m-1
        } else {
            has = m
            break
        }
    }

    elapsed = time.Since(start)
    fmt.Println("logn -> ",elapsed, "has -> ", has)
}