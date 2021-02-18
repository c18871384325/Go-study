package main

import "fmt"

// 包级别的变量定义
var packagevar string = "package hello" //（在包下的函数生效）

func main() {
	// var 定义变量，msg 为定义的变量名， string 表示变量为字符串类型，hello word 为变量值
	var msg string = "hello"

	// 函数级别的变量定义 (仅在函数内生效)
	//变量的定义，当在包级别和函数快级别同时定义相同名的变量，当在函数内子快引用时，函数快内定义的变量生效，子块内查找变量时会先查找自己块内的，再查找父快内的变量，查找到后不会再往上层查找该变量
	var funcvar string = "func hello"
	{
		// 函数内{ } 的方括号包裹的是快，快级别的变量作用于快内
		// 注意：子快内定义的变量无法作用于父快，即无法作用于上层快，上层快包括函数及包级别的变量可以作用于子快内
		var block string = "block hello"
		fmt.Println(packagevar, msg, block)
	}
	fmt.Println(msg, packagevar, funcvar)
}





