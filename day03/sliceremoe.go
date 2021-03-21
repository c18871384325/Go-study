package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	copy(slice[1:], slice[2:])
	fmt.Println(slice)
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
