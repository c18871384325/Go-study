package main

import (
	"fmt"
	"time"
)

type task struct {
	id        int
	name      string
	starttime time.Time
	endtime   time.Time
	status    int
	user      string
}


func changetask(ptask task) {
	ptask = task{
		id: 5,
		name: "cisco",
		user: "cisco",
	}
	fmt.Println(ptask)
}


func main() {
	// type task int    // 定义一个结构体task ，结构体类型为 int
	// var cnt task
	// fmt.Printf("%T, %#v", cnt, cnt)
	var demo task

	demo = task{1, "tencent", time.Now(), time.Now().Add(24 * time.Hour), 1, "huawei"}  // 定义值的顺序必须与结构体内属性的定义相同，都进行赋值
	fmt.Printf("%T, %#v\n", demo, demo)

	// 第二种赋值方式，为指定的属性赋值
	demo = task{id: 2, name: "huawei", status: 5}  // 未赋值的属性为对应属性类型的零值
	fmt.Println(demo)

	demo = task{
		id: 3,
		name: "tencent",
		starttime: time.Now().Add(24 * time.Hour),   // 注意，最后一个逗号不能省略，否则会报错，如果大括号跟最后一个在同一行，则不用逗号
	}
	fmt.Println(demo)


	// 结构体属于值类型，函数类赋值不会有影响

	changetask(demo)  // 将demo 传递给函数，在函数内部修改值，对外部的dmeo不影响，结构体是值类型，非引用类型，引用类型会受影响
	fmt.Println(demo)


	// 指针类型的结构体
	var ptask *task  
	fmt.Printf("%T\n", ptask)
	fmt.Printf("%#v\n", ptask)
	// (*main.task)(nil) ， 类型为  *main.task ,零值为 nil

	ptask = &task{   // 解引用为ptask赋值
		id: 9,
		name: "h3c",
		user: "h3c",
	}
	fmt.Println(ptask)


	// 单个结构体赋值及引用
	ptask.starttime = time.Now()
	fmt.Println(ptask.starttime)


}
