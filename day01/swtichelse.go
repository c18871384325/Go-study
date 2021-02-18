package main

import "fmt"

func main() {
	// 90分以上，A  80分以上B， 60分以上C， 60分以下D

	var score float32

	fmt.Print("请输入分数：")
	fmt.Scan(&score)
	fmt.Println("你输入的分数是：", score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
	// if 必须有， else if  可以有0-N个， else 只能有1个
	// if score >= 90 {
	// 	fmt.Println("A")
	// } else if score >= 80 {
	// 	fmt.Println("B")
	// } else if score >= 60 {
	// 	fmt.Println("C")
	// } else {
	// 	fmt.Println("D")
	// }
}
