package main

import (
	"fmt"
	"time"
)
 
func main() {
	// 命名结构体，结构体内嵌套结构体

	type User struct{
		name string
		addr string
		tel string
	}

	type task struct{
		id int
		taskname string
		starttime time.Time
		endtime time.Time
		user User    //  将属性 user 类型定义为结构体 User 类型
	}

	// 零值为对应属性类型的零值
	todotask := task{}

	fmt.Printf("%T\n", todotask)
	fmt.Printf("%#v\n", todotask)
	// struct { id int; name string; age int }{id:1, name:"kobi", age:35}

	
	// 赋值
	todotask = task{
		id: 1,
		taskname: "坐地铁",
		starttime: time.Now(),
		endtime: time.Now().Add(1 * time.Hour),
		user: User{
			name: "kobi",
			addr: "america",
			tel: "123456",
		},   // 注意：这里的逗号不能省略
	}
	fmt.Printf("%#v\n", todotask)


	// 单个属性赋值
	todotask.user.name = "email"
	fmt.Println(todotask.user.name)

}