package main

import (
	"fmt"
)

type Task struct {
	id   int
	name string
	tel  int
}

func (task *Task) SetName(name string) {
	task.name = name
}

func (task *Task) GetName() string {
	return task.name
}

func main() {
	// 为函数定义接收者为结构体类型 Task ,变量名为 task，并使用指针接收

	task := &Task{name: "beego"} // 定义变量为指针类型的Task 结构体
	task.SetName("kobi")         // 因为是指针传递，此处可以修改结构体 task 的值
	fmt.Println(task.GetName())
	// kobi

}
