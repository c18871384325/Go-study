package main

import (
	"fmt"
	"strconv"
)

// strconv ，字符串转换包，该包可以进行 字符串，浮点数，整形之间的转换

func main() {
	var (
		intval = 65
		floatval = 1.3
		strval = "3.3"
	)

	v, err := strconv.Atoi(strval)  // 会打印一个值和错误信息
	fmt.Println(v, err)

	fmt.Println(intval, floatval, strval)
	fmt.Println(strconv.FormatFloat(floatval, 'f', 10, 64))   // 将float转换为字符串格式
}

