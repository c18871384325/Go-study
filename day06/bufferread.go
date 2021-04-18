package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	read1 := bytes.NewReader([]byte("this is demo"))   // 初始化read1

	ctx := make([]byte, 3)
	n, _ := read1.Read(ctx)
	fmt.Println(string(ctx[:n]))    // 将read1内容读取到 ctx 字节切片

	read1.Reset([]byte("qing"))  // 重置read1的内容

	read1.WriteTo(os.Stdout)



}