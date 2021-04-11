package main

import (
	"fmt"
)

type Task struct {
	id   int
	name string
	tel  int
}

func (task Task) SetName(name string) {
	task.name = name
}

func (task Task) GetName() string {
	return task.name
}

func main() {
	// 定义结构体，在函数指定接收者为结构体

	task := Task{name: "beego"}
	task.SetName("huawei") // 注意，SetName 函数并不会修改task.name, 因为结构体为值传递，此处的调用函数相当于把task.name的值拷贝到函数内，在函数内进行的修改并不影响函数外部的值，因此修改不生效，可以用指针传递的方式使得值生效
	fmt.Println(task.GetName())
	// beego
}
