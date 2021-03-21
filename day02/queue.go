package main

import "fmt"

func main() {
	// 队列
	// 先进先出
	queue := []string{}

	queue = append(queue, "a", "b")   // 使用append向切片内按顺序向元素写入值，
	queue = append(queue, "c")

	// pop
	x := queue[0]
	queue = queue[1:]    // 删除切片第一个元素
	fmt.Println("1", x)

	x = queue[0]
	queue = queue[1:]
	fmt.Println("2", x)

	x = queue[0]
	queue = queue[1:]
	fmt.Println("2", x)
}