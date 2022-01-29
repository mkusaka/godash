package main

import (
	"fmt"
	"github.com/mkusaka/godash"
)

func main() {
	fmt.Println("hello world")
	org := []int{1, 2, 3, 4}
	filterd = godash.Filter(org, func(e int) bool {
		return e%2 == 0
	})
	fmt.Printf("filterd : %+v", filtered)
}
