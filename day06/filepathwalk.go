package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, file os.FileInfo, err error) error {
		fmt.Println(path, file.Size(), file.IsDir())
		return nil
	})

	// filepath.walk ， 获取指定目录或文件的信息， "."  表示访问路径， path 为文件当前路径名， file ,表示将每个文件的fileinfo信息传递给file，file可以使用函数查看文件信息，如 file.Size(),file.Name(),  err 可以返回bool型的值

}
