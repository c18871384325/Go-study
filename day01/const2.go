package main

import (

	"fmt"
)

func main() {

	// const B   不允许重复定义常量

	const (
		A = "test"
		B  // 使用前一个常量的值进行初始化（B => A）
		C // 使用前一个常量的值进行初始化（C => B）
		D = "testD"
		E // 使用前一个常量的值进行初始化（E => D）
		F
	)
	fmt.Println(B, C)
	fmt.Println(E, F)
}