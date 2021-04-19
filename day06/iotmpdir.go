package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	tmp, err := ioutil.TempDir("./test", "log")   // 在 ./test 目录下生成一个 beggo 开头的随机字符串后缀目录
	if err != nil {
		return
	}
	file, err := os.Create(filepath.Join(tmp, "1.log"))   // 根据上面创建的目录，在目录下创建 1.log 文件
	file.WriteString(time.Now().Format("15:04:05"))  // 向创建的文件写入内容
	file.Close()
	fmt.Println(os.RemoveAll(tmp))   // 移除目录及目录下的所有文件
	
	
}