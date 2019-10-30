package main

import (
	"fmt"
)

func main() {
	sum := 0
	n := 1000000000
	fmt.Println("n =", n)
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Println("sum =", sum)
}
