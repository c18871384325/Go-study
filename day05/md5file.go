package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// 计算文件的md5值，用户通过 -p 选项指定文件路径，返回文件的md5 值

	var filepath string
	flag.StringVar(&filepath, "p", "", "文件路径")
	flag.Parse()
	if filepath == "" {
		return
	}

	file, err := os.Open(filepath)

	defer file.Close()

	if err != nil {
		fmt.Println(nil)
		return
	}

	hasher := md5.New()

	ctx := make([]byte, 1024)
	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		hasher.Write(ctx[:n])
	}
	fmt.Printf("%x", hasher.Sum(nil))
}
