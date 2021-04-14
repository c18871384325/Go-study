package main

import (
	"fmt"
	"os"
)

func main() {

	os.MkdirAll("password/a/b/c", 0644)
	dir, err := os.Open("D:\\src\\Go-study\\day05")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1) // -1 表示返回所有目录下的文件名
	fmt.Println(names)
	// [anonymous.go dir.go dirname.go filedemo fileexsits.go fileinfo.go fileseek.go md5file.go method.go password ptask.go readfile.go recoverlog.go task.go taskv1.go taskv2.go taskv3.go taskv4.go taskv5.go writefile.go]

}
