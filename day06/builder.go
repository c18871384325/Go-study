package main

import (
	"fmt"
	"strings"
)

func main() {
	var build strings.Builder   // 定义一个内存结构体指针，通过build可以向其中写入内容，内容是直接写入内存
	build.WriteByte('m')        // 写入字节类型
	build.WriteRune('中')        // 写入Rune类型
	build.WriteString("国人")     // 写入字符串类型
	fmt.Println(build.String()) // m中国人   , 将内存写入的数据输出为字符串

	fmt.Println(build.Len()) // 读取内存中build的长度
	build.Reset()            // 清空内存中build结构体的内容

	fmt.Println(build.String())
	fmt.Println(build.Len()) // 清空后发现长度为0

}
