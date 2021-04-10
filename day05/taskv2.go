package main

import (
	"fmt"
)

func main() {
	// 匿名结构体，结构体内嵌套结构体, 属性内部的结构体值继承上层结构体的值

	type User struct {
		name string
		addr string
		tel  string
	}

	type task struct {
		id   int
		name string
		num  int
		User //  // 直接将结构体类型作为结构体的属性，也有属性名称（User）
	}

	// 赋值
	todotask := task{
		id:   1,
		name: "study",
		num:  10,
		User: User{
			name: "default",
			addr: "tencent",
		},
	}
	fmt.Println(todotask)

	fmt.Println(todotask.name) // 当嵌套为匿名结构体时，结构体内与嵌套结构体内存在相同的属性，打印结果为上层结构体属性值
	fmt.Println(todotask.User.name)
	fmt.Println(todotask.addr) // 当嵌套为匿名结构体时，可以直接使用嵌套结构体内的属性打印
	// study
	// default
	// tencent

	// 单个结构体赋值
	todotask.User.addr = "qianhai"
	fmt.Println(todotask.User.addr)

}
