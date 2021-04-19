package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	ctx, err := ioutil.ReadFile("iocopy.go")    // ioutil.ReadFile  ，直接打开一个文件并读取到 ctx  
	if err != nil {
		return
	}
	fmt.Println(string(ctx))
}