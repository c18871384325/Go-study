package main

import "fmt"

func main() {

	// 如果有卖西瓜的，就买1个包子，否则买10个包子

	var y string

	fmt.Print("有卖西瓜的吗：")
	fmt.Scan(&y)

	switch y { //  switch 等同于 if  case 为条件判断，成立则执行case下的代码，default为case条件都不满足时最后执行的代码（等同于else）
	case "yes", "y", "Y": // switch 方式支持同时判断多个
		fmt.Println("买一个包子")
	case "no", "N", "n":
		fmt.Println("买十个包子")
	default:
		fmt.Println("输入错误")
	}
	// if y == "yes" {
	// 	fmt.Println("买1个包子")
	// } else {
	// 	fmt.Println("买十个包子")
	// }
}
