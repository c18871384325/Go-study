package main

import "fmt"

func main() {
	name := ""
	fmt.Print("请输入你的名字：")

	fmt.Scan(&name) // Scan ，会读取标准输入的值，将输入的结果写入 name 变量的内存指针，因此之后print打印name变量值，值为输入的值

	fmt.Println("你输入的名字是：", name)

	age := 10 // int类型的变量

	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age) // 注意：age变量是int类型的变量，此处标准输入需要输入int类型的值，否则无法赋值

	msg := ""

	fmt.Print("请输入你的信息：")
	fmt.Scan(&msg)

	fmt.Println(msg)
}
