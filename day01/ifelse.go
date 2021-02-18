package main

import "fmt"

func main() {

	// 如果有卖西瓜的，就买1个包子，否则买10个包子

	var y string

	fmt.Print("有卖西瓜的吗：")
	fmt.Scan(&y)

	if y == "yes" {
		fmt.Println("买1个包子")
	} else {
		fmt.Println("买十个包子")
	}
}
