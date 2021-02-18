package main 

import "fmt"

func main() {
	// 打印99乘法表
	fmt.Println("1 * 1 = 1")
	fmt.Println("1 * 2 = 2   2 * 2 = 4" )
	fmt.Println("1 * 3 = 3   2 * 3 = 6  3 * 3 = 9")
	
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d   ", i, j, i * j)
		}
		fmt.Println("")
	}
}