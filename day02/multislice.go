package main

import "fmt"

func main() {
	// 切片里面嵌套切片
	multislice := [][]string{}

	// 嵌套切片的赋值
	multislice = append(multislice, []string{0: "huwei", 1: "cisco"})
	multislice = append(multislice, []string{0: "jzt", 1: "szse"})
	fmt.Printf("%T, %#v\n", multislice, multislice)  // [][]string, [][]string{[]string{"huwei", "cisco"}, []string{"jzt", "szse"}}


	// 嵌套切片的遍历
	for i := 0; i < len(multislice); i++ {
		fmt.Println(i, multislice[i])
		// 0 [huwei cisco]
		// 1 [jzt szse]
	}


	// 嵌套切片的第二种赋值方式   修改
	multislice[0][1] = "h3c"
	fmt.Println(multislice)   // [[huwei h3c] [jzt szse]]


	// 嵌套切片的第三种赋值方式  增加
	multislice[1] = append(multislice[1], "hd" )
	fmt.Println(multislice)  // [[huwei h3c] [jzt szse hd]]
}