package main

import (
	"fmt"
)


func main() {
	name, desc := "czhe", 26
	func(name string) {
		desc = 27
		fmt.Println(name,desc)
	}("xqin")
	fmt.Println(name, desc)
}