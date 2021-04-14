package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logfile, err := os.OpenFile("filedemo", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer logfile.Close()

	log.SetOutput(logfile) // 将日志的标准输出写入到logfile，实际就是写入到打开的文件 filedemo内，此处

	log.Println("This is logfile write")
	log.Println("demo")
	// 	2021/04/13 09:16:14 This is logfile write
	// 2021/04/13 09:16:14 demo
}
