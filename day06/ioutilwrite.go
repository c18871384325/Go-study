package main

import (
	//"fmt"
	"io/ioutil"
	"os"
)

func main() {
	ioutil.WriteFile("./test/util.go", []byte("this is ioutilwrite"), os.ModePerm) // 直接创建文件 ./test/util.go 并写入指定内容
}
