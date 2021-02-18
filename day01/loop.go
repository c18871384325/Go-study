package main

import "fmt"

func main() {
	var index = 0
	var total = 0

	for {   // 不带任何条件的for语句是死循环，一直循环下去
		total += index
		index++
		if index > 100 {
			break   // 当循环到index大于100时，退出for循环
		}
	}   // 该for循环相当于计算 1+..100 的值  

	fmt.Println(total)
}