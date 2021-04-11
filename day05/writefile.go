package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("kobi.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm) // os.ModePerm 为文件权限，默认为 0777  ， 可以自己指定值为 0700 0644 等
	// os.O_CREATE  为当文件不存在时创建文件
	// os.O_APPEND  文件存在时从后面追加写入  // 与 trunc 不同时存在
	// os.O_RDWR 读写方式打开文件
	// os.O_TRUNC  清空文件内容后写入
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write([]byte("demo\n")) // 写入文件时  \n 换行符
	file.Write([]byte("file\t")) // 写入文件时  \t 制表符
	file.Write([]byte("123"))

}
