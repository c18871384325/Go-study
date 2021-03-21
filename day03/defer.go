package main

import(
	"fmt"
)

func main() {
	defer func(){
		recover()
	}()
	fmt.Println("befover")
	panic("after error")
	fmt.Println("after")
}