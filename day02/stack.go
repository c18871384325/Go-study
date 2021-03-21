package main

import "fmt"

func main() {
	// 堆栈
	// 先进后出

	// push 
	// append
	stack := []string{}
	stack = append(stack, "a")
	stack = append(stack, "b")
	stack = append(stack, "c")


	//pop
	// 后面移除

	x := stack[len(stack)-1:len(stack)]
	stack = stack[:len(stack)-1]
	fmt.Println(x)

	x = stack[len(stack)-1:len(stack)]
	stack = stack[:len(stack)-1]
	fmt.Println(x)

	x = stack[len(stack)-1:len(stack)]
	stack = stack[:len(stack)-1]
	fmt.Println(x)


}