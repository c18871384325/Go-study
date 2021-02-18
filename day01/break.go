package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // break 表示退出当前for循环代码
		}
		fmt.Println(i)
	}
}
