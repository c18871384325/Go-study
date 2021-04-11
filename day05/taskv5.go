package main

import (
	"fmt"
)

type Task struct {
	id   int
	name string
	tel  int
}

func (task *Task) SetTel(tel int) {
	task.tel = tel
}

func (task2 Task) SetTel2(tel int) {
	task2.tel = tel
}

func main() {

	// 定义了两个函数接收者，一个为指针类型，一个为值类型  ，值是否可以修改，主要看接收者类型，接收者为指针则可以修改，值类型则不可以修改

	task := Task{tel: 123456} // 使用值类型传递
	task.SetTel(789)          //  此处可以修改tel的值，虽然 task定义非 &Task 指针类型，但是Go语言的语法糖在编译时会将结构体类型转换为指针类型，因此此处可以修改成功
	fmt.Printf("%#v\n", task)
	// main.Task{id:0, name:"", tel:789}

	task2 := &Task{tel: 123456}
	task2.SetTel2(789) // 此处修改不成功，虽然task2 定义类型为指针，但是编译时Go语言的语法糖会将指针类型修改为值传递类型，因为函数的接收者为值传递类型，因此此处无法修改
	fmt.Printf("%#v\n", task2)
	// &main.Task{id:0, name:"", tel:123456}

}
