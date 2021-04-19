package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".") // files 获取到的内容是指定目录下的文件或文件夹的 os.stat 状态，可以通过os.stat的函数查看详细内容
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("D", file.Size(), file.Name(), os.ModePerm)
		} else {
			fmt.Println("F", file.Size(), file.Name(), os.ModePerm)
		}
	}
}
