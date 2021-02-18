package main

import "fmt"

func main() {
	// 买10个包子，如果遇到卖西瓜的，买一个西瓜

	fmt.Println("买10个包子")

	var y string
	fmt.Print("有没有卖西瓜的：")
	fmt.Scan(&y)

	if y == "yes" {
		fmt.Println("买一个西瓜")
	}
}
