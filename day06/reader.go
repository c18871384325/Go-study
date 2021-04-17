package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("!QAZ@WSX3edc4rfv") // 生成一个新的 reader 结构体指针存储数据
	ctx := make([]byte, 5)
	for {
		n, err := reader.Read(ctx) //   从reader对象读取数据到ctx
		if err != nil {
			break
		}
		fmt.Println(string(ctx[:n])) // 打印
	}
	reader.Seek(0, os.SEEK_SET) // 上面读取完成后光标处于最后了，无法再读取，通过seek将光标移动到最前面，可以重新读取
	n, err := reader.Read(ctx)
	fmt.Println(n, err, string(ctx[:n]))
	fmt.Println(reader.Size()) // 获取reader结构体指针的长度

	reader.Reset("123456789") // 重新为 reader 结构体赋值数据
	reader.WriteTo(os.Stdout) // 将reader的内容写入到io
}
