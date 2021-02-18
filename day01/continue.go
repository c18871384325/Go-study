package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // continue 表示跳过当次循环，执行下次循环判断
		}
		fmt.Println(i)
	}
}
