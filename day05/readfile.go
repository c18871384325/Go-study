package main

import (
	"fmt"
	"os"
	"io"
)

func main() {

	// Go 标准的一个操作文件的流程
	file, err := os.Open("filedemo")
	// fmt.Println(file, err)    // &{0xc0000ce780} <nil>     文件正常打开时，err为 nil
	if err != nil {
		return
	}
	defer file.Close()        // 不管是否有错误，函数退出时都关闭文件
	ctx := make([]byte, 1024) // 定义ctx 为byte类型的切片，大小为 1024byte ， 1K
	for {   // 文件读取结束后，会有EOF的err信息，一般可以用EOF退出
		n, err := file.Read(ctx) // 将open 打开的文件内容 通过 file.Read 读取到 ctx 切片内   ,   Read 函数会返回2个值， n 为文件内容的长度，单位byte，err 为错误信息
		if err == io.EOF {
			break
		}
		fmt.Println(n, err, string(ctx), len(string(ctx[:n])))
		//
	}
}


