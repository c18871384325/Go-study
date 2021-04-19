package main

import (
	//"fmt"
	"io/ioutil"
	"time"
)

func main() {
	tmpfiles, _ := ioutil.TempFile("./test", "log")   // 在 ./test 目录下创建以 log 开头的随机字符串后缀结尾的文临时文件
	tmpfiles.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))    // 向临时创建的文件写入内容

}