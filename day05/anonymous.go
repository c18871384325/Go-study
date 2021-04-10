package main

import (
	"fmt"
)

// 匿名结构体，直接将变量定义为结构体，并不使用 type 方式定义一个结构体


func main() {
	var user struct{   // 直接将变量定义为结构体类型
		id int
		name string
		age int
	}

	// 赋值
	user = struct{id int; name string; age int}{1, "kobi", 35}

	fmt.Printf("%T\n", user)
	fmt.Printf("%#v\n", user)
	

	// 给结构体某个属性赋值
	user.name = "beego"
	fmt.Println(user.name)
	fmt.Println(user)

}