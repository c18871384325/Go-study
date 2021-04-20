package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // 生成一个新的bufio扫描，从标准输入读取，也可以从打开的文件流读取
	fmt.Print("请输入内容： ")
	scanner.Scan() // 扫描一个新的输入，读取标准输入的内容，遇到Enter键，读取完成
	fmt.Println("输入的内容是：", scanner.Text())

	// scan 文件扫描完会返回一个false ,可以用for循环扫描
}
