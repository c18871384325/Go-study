package main

import (
	"fmt"
	"os"
)

// 判断文件是否存在

func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func main() {
	fmt.Println(FileIsExists("filedemo.txt"))

}
