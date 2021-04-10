package main

import (
	"fmt"
)

func main() {

	type User struct {
		name string
		addr string
		tel  string
	}

	type task struct {
		id    int
		name  string
		num   int
		*User // 将结构体指针作为外层结构体的属性，零值为nil
	}

	// 赋值：
	ptask := task{
		id:   1,
		name: "beego",
		num:  2,
		User: &User{}, // 初始化指针
	}
	fmt.Println(ptask)

	ptask.User = &User{
		name: "ptask1",
		addr: "sz",
	}
	fmt.Printf("%#v\n", ptask.User)
	// &main.User{name:"ptask1", addr:"sz", tel:""}

	// 当定义的结构体指针被多个变量使用时，一个修改，另外一个的值也随着修改
	ptaskv1 := ptask

	ptaskv1.name = "kobi"
	fmt.Println(ptask)
	fmt.Println(ptaskv1)

	ptaskv1.User.name = "kobi" // 当指针类型的结构体值被修改，其他的引用值也会被修改
	fmt.Println(ptask.User)
	fmt.Println(ptaskv1.User)
	// &{kobi sz }
	// &{kobi sz }

}
