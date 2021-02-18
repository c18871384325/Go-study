package main

import "fmt"


var huawei string = "cisco" // 在包级别定义的变量，可以在代码中不使用

func main() {
	// 函数内包括快内定义的变量，不使用会报错，包级别变量无该项限制
	var name string = "shinet"
	// 定义了变量类型，并且初始化了变量值为 shitnet

	var zeroString string 
	// 定义变量类型，但不初始化值
	// 初始化值使用类型对应的零值（默认字符串空值""）

	var typeString = "shient"
	// 定义变量省略类型，通过对应的值推导出变量的值

	shortString := "shiNet"
	// 短变量声明，必须在函数内包含函数子快内使用，通过对应值推导出类型

	fmt.Println(name, zeroString, typeString, shortString)

	typeString = "zabbix"  // 更新已定义的变量的值（赋值）

	fmt.Println(typeString)

}