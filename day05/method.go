package main

import (
	"fmt"
)

type Task struct {
	id   int
	name string
	user string
}

func (task Task) SetUser(user string) {
	task.user = user
}

func (task *Task) SetName(name string) {
	task.name = name
}

func main() {

	task := Task{}
	methode1 := Task.SetUser // 将函数接收者跟函数名定义为一个方法
	fmt.Printf("%T\n", methode1)
	// func(main.Task, string)   methode1 的类型为一个类似函数类型，接受2个参数，一个Task类型的变量，一个值
	fmt.Println(task)

	methode1(task, "huawei") // 无法修改，因为SetUser 类型为值传递
	fmt.Printf("%#v\n", task)
	// main.Task{id:0, name:"", user:""}

	task2 := &Task{}
	methode2 := (*Task).SetName  // 注意：SetName 函数为指针接收者，因此此处 Task 传参必须为指针，否则会报错
	fmt.Printf("%T\n", methode2) // func(*main.Task, string)

	methode2(task2, "cisco")   // 可以修改， SetName 函数接收者为指针类型，此处 methode2 调用为指针类型，可以修改
	fmt.Printf("%#v\n", task2) // &main.Task{id:0, name:"cisco", user:""}

}
