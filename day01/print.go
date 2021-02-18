package main 

import "fmt"

func main() {
	var name = "myname"

	fmt.Println(name)  // 打印name 变量并换行
	fmt.Println("*")
	fmt.Print(name)  // 打印name变量，不换行直接打印下一行
	fmt.Print("*")
	fmt.Printf("%T, %v, %#v", name, name, name)
		// printf , 提供一个占位符， T：变量类型， v：打印变量值， #v: 为变量值加双引号打印
}