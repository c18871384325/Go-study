package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
	"strings"
)

func main() {
	file, err := os.Open("./test/demo")
	if err != nil {
		return
	}
	defer file.Close()
	buffer := bytes.NewBuffer([]byte(""))

	//将打开的file内容copy到buffer内
	_, err = io.Copy(buffer, file)
	if err != nil {
		return
	}

	// 打印buffer的内容，发现文件内容已写入buffer
	fmt.Println(buffer.String())   // abcde12345  
	file.Close()


	file1, err := os.Open("./test/demo")
	buffer1 := &strings.Builder{}    

	ctx := make([]byte, 1024)   // 定义buffer 
	_, _  = io.CopyBuffer(buffer1, file1, ctx)    // copy 时指定buffer的大小
 	fmt.Println(buffer1.String())

}	