package main

import "fmt"

func main() {
	const (
		A = iota // 在常量内使用 iota 初始化值为0, 每次调用+1
		B // 值为1
		C // 值为2
		D // 值为3
		E // 值为4
	)
	fmt.Println(A,B,C,D,E)
}
