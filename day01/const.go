package main

import "fmt"

//常量的合并定义
const (
	packageName string = "main"
	packageMsg = "error"
)	

func main() {
	// 常量的值定义后无法修改
	const name string = "myname" // 声明常量 name，并赋值为 myname
	const shinet = "cisco" // 常量声明定义后，可以不使用
	fmt.Println(name)
}